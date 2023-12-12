-- MySQL database script

-- Check if database exists, and create if it does not
CREATE DATABASE IF NOT EXISTS wgpage;
USE wgpage;

-- Table: shopping_items
CREATE TABLE IF NOT EXISTS shopping_items (
    id INT AUTO_INCREMENT,
    item VARCHAR(100),
    bought BOOLEAN,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- MySQL script complete