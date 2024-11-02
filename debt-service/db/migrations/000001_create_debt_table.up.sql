CREATE TABLE IF NOT EXISTS debts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    jshshir INT NOT NULL,
    address TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    price_paid DOUBLE PRECISION NOT NULL,
    acquaintance TEXT,
    collateral TEXT,
    deadline TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);