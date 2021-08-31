CREATE TABLE IF NOT EXISTS payment_codes(
   id uuid PRIMARY KEY,
   payment_code TEXT,
   name TEXT,
   status TEXT,
   expiration_date timestamptz,
   created_at timestamptz,
   updated_at timestamptz
);
