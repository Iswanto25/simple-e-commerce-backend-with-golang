-- migrations/20240705131708_create_table_users.up.sql

CREATE TABLE users (
    id VARCHAR(15) PRIMARY KEY,
    photos VARCHAR(25) ,
    name VARCHAR(255) ,
    username VARCHAR(255) ,
    email VARCHAR(255) ,
    password VARCHAR(255) ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);