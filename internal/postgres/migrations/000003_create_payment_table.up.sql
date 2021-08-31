CREATE TABLE IF NOT EXISTS payments(
   id uuid PRIMARY KEY,
   payment_code TEXT,
   transaction_id TEXT,
   name TEXT,
   amount INTEGER,
   created_at timestamptz,
   updated_at timestamptz
);
