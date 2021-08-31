CREATE TABLE IF NOT EXISTS inquiries(
   id uuid PRIMARY KEY,
   payment_code TEXT,
   transaction_id TEXT,
   created_at timestamptz,
   updated_at timestamptz
);
