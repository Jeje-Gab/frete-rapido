CREATE SCHEMA IF NOT EXISTS frete;

CREATE TABLE IF NOT EXISTS frete.quote_requests (
    id SERIAL PRIMARY KEY,
    recipient_zipcode VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS frete.quotes (
    id SERIAL PRIMARY KEY,
    quote_request_id INT REFERENCES frete.quote_requests(id),
    carrier_name TEXT NOT NULL,
    service TEXT,
    deadline INT,
    price NUMERIC,
    created_at TIMESTAMP DEFAULT NOW()
);
