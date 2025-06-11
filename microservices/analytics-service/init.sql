-- init.sql

CREATE TABLE IF NOT EXISTS order_analytics (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
