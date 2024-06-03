CREATE DATABASE IF NOT EXISTS gopay;

USE gopay;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_merchant BOOLEAN NOT NULL,
    balance FLOAT DEFAULT 0
);