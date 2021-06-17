BEGIN;
CREATE TABLE IF NOT EXISTS paymentcode(
    id varchar(255) NOT NULL PRIMARY KEY,
    payment_code varchar(30) NOT NULL,
    name varchar(100) NOT NULL,
    status varchar(20) NOT NULL,
    expiration_date timestamptz NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
    );

COMMIT;
