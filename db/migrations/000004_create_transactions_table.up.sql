CREATE TABLE IF NOT EXISTS transactions (
    id CHAR(36) PRIMARY KEY,
    order_id CHAR(36) NOT NULL,
    transaction_method VARCHAR(60) NOT NULL,
    amount DECIMAL(6,2) NOT NULL,
    status int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);