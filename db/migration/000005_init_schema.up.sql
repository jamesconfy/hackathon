ALTER TABLE deposits ADD COLUMN account_id uuid;

ALTER TABLE deposits ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE;