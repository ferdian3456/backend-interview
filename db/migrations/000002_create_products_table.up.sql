CREATE TABLE IF NOT EXISTS products (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    name VARCHAR(127) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(6,2) NOT NULL,
    available VARCHAR(9) DEFAULT 'Ready',
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);