CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal', 'transfer');

CREATE TABLE transactions
(
  id                     uuid PRIMARY KEY          DEFAULT uuid_generate_v4(),
  type                   transaction_type NOT NULL,
  amount                 NUMERIC(10, 2)   NOT NULL,
  source_account_id      INTEGER REFERENCES accounts (id),
  destination_account_id INTEGER REFERENCES accounts (id),
  created_at             TIMESTAMP        NOT NULL DEFAULT NOW(),
  is_rolled_back         BOOLEAN          NOT NULL DEFAULT FALSE
);