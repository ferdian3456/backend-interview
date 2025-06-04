CREATE TABLE IF NOT EXISTS orders (
    id CHAR(36) PRIMARY KEY,
    product_id CHAR(36) NOT NULL,
    customer_id CHAR(36) NOT NULL,
    seller_id CHAR(36) NOT NULL,
    created_by VARCHAR(22) NOT NULL,
    amount DECIMAL(6,2) NOT NULL,
    status int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES users(id) ON DELETE CASCADE
);