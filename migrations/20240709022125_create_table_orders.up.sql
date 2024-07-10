-- migrations/20240709022125_create_table_orders.up.sql

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id VARCHAR REFERENCES products(id),
    idCustomer VARCHAR REFERENCES customers(id),
    quantity INT,
    total DECIMAL(10, 2),
    status VARCHAR(255) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
