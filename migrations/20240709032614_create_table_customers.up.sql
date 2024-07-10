-- migrations/20240709032614_create_table_customers.up.sql

CREATE TABLE customers (
    id VARCHAR(15) PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
