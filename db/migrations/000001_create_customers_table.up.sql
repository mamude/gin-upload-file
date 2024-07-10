CREATE TABLE IF NOT EXISTS customers(
    id SERIAL PRIMARY KEY,
    cpf VARCHAR (14) NOT NULL,
    private INT NOT NULL,
    incomplete INT NOT NULL,
    date_last_purchase TIMESTAMP,
    average_ticket DECIMAL(10,2) DEFAULT(0.00),
    last_purchase_ticket DECIMAL(10,2) DEFAULT(0.00),
    most_frequent_store VARCHAR(18),
    last_purchase_store VARCHAR(18),
    created_at TIMESTAMP DEFAULT 'now()'
);