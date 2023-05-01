package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/escalopa/fingo/wallet/internal/adapters/locker"

	"github.com/escalopa/fingo/pb"
	pkgdb "github.com/escalopa/fingo/pkg/db"
	"github.com/escalopa/fingo/pkg/tracer"
	"github.com/escalopa/fingo/wallet/internal/adapters/db"
	mygrpc "github.com/escalopa/fingo/wallet/internal/adapters/grpc"
	"github.com/escalopa/fingo/wallet/internal/adapters/numgen"
	"github.com/escalopa/fingo/wallet/internal/application"
	"github.com/lordvidex/errs"

	pkgerror "github.com/escalopa/fingo/pkg/error"
	grpctls "github.com/escalopa/fingo/pkg/tls"
	pkgtracer "github.com/escalopa/fingo/pkg/tracer"

	pkgvalidator "github.com/escalopa/fingo/pkg/validator"
	"github.com/escalopa/goconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := goconfig.New()
	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create validator
	v := pkgvalidator.NewValidator()
	log.Println("validator created")

	// Create database connection
	conn, err := pkgdb.New(c.Get("WALLET_DATABASE_URL"))
	pkgerror.CheckError(err, "failed to create database connection")
	log.Print("database connection created")

	// Migrate database
	pkgerror.CheckError(pkgdb.Migrate(conn, c.Get("WALLET_DATABASE_MIGRATION_PATH")), "failed to migrate database")
	log.Print("database migrated")

	ur := db.NewUserRepository(conn)
	cr := db.NewCardRepository(conn)
	ar := db.NewAccountRepository(conn)
	tr := db.NewTransactionRepository(conn)

	// Create a new number generator
	cardNumberLength, err := strconv.Atoi(c.Get("WALLET_CARD_NUMBER_LENGTH"))
	pkgerror.CheckError(err, "failed to convert WALLET_CARD_NUMBER_LENGTH to int")
	cng := numgen.NewNumGen(cardNumberLength)
	log.Println("Card number generator created with card number length:", cardNumberLength)

	// Create a new tracer
	t, err := pkgtracer.LoadTracer(
		c.Get("WALLET_TRACING_ENABLE"),
		c.Get("WALLET_TRACING_JAEGER_ENABLE"),
		c.Get("WALLET_TRACING_JAEGER_AGENT_URL"),
		c.Get("WALLET_TRACING_JAEGER_SERVICE_NAME"),
		c.Get("WALLET_TRACING_JAEGER_ENVIRONMENT"),
	)
	pkgerror.CheckError(err, "failed to load tracer")
	tracer.SetTracer(t)
	log.Println("tracer created")

	// Create an ids locker
	cleanupDuration, err := time.ParseDuration(c.Get("WALLET_LOCKER_CLEANUP_DURATION"))
	pkgerror.CheckError(err, "failed to convert WALLET_LOCKER_CLEANUP_DURATION to time.Duration")
	l := locker.NewLocker(appCtx, cleanupDuration)

	// Create use cases
	uc := application.NewUseCases(
		application.WithValidator(v),
		application.WithLocker(l),
		application.WithUserRepository(ur),
		application.WithCardRepository(cr),
		application.WithAccountRepository(ar),
		application.WithTransactionRepository(tr),
		application.WithCardNumberGenerator(cng),
	)

	var opts []grpc.ServerOption
	// Load TLS certificates
	err = loadTls(c, &opts)
	pkgerror.CheckError(err, "failed to load wallet TLS certificates")

	// Load auth interceptor
	err = loadInterceptor(c, &opts)
	pkgerror.CheckError(err, "failed to load auth interceptor")

	// Start gRPC server
	pkgerror.CheckError(start(c, uc, opts), "failed to start gRPC server")
}

func start(c *goconfig.Config, uc *application.UseCases, opts []grpc.ServerOption) error {
	// Create a gRPC server object
	handler := mygrpc.NewWalletHandler(uc)
	grpcS := grpc.NewServer(opts...)
	pb.RegisterWalletServiceServer(grpcS, handler)
	reflection.Register(grpcS)

	// Start the server
	port := c.Get("WALLET_GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return errs.B(err).Msg(fmt.Sprintf("failed to listen on port %s", port)).Err()
	}
	log.Println("starting gRPC server on port", port)
	err = grpcS.Serve(lis)
	if err != nil {
		return errs.B(err).Msg(fmt.Sprintf("failed to serve on port %s", port)).Err()
	}
	return nil
}

func loadTls(c *goconfig.Config, opts *[]grpc.ServerOption) error {
	// Enable TLS if required
	creds, err := grpctls.LoadServerTLS(
		c.Get("WALLET_GRPC_TLS_ENABLE"),
		c.Get("WALLET_GRPC_TLS_CERT_FILE"),
		c.Get("WALLET_GRPC_TLS_KEY_FILE"),
	)
	if err != nil {
		return err
	}
	*opts = append(*opts, grpc.Creds(creds))
	return nil
}

func loadInterceptor(c *goconfig.Config, opts *[]grpc.ServerOption) error {
	creds, err := grpctls.LoadClientTLS(
		c.Get("TOKEN_GRPC_TLS_ENABLE"),
		c.Get("WALLET_TOKEN_GRPC_TLS_USER_CERT_FILE"),
	)
	if err != nil {
		return err
	}
	interceptor, err := mygrpc.NewAuthInterceptor(c.Get("TOKEN_GRPC_URL"), creds)
	if err != nil {
		return errs.B(err).Msg("failed to create token gRPC interceptor").Err()
	}
	*opts = append(*opts, grpc.UnaryInterceptor(interceptor.Unary()))
	log.Println("created gRPC token interceptor")
	return nil
}
