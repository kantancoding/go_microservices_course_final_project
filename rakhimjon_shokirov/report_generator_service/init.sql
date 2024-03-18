CREATE USER 'report_generator_user'@'%' IDENTIFIED BY 'Auth123';

CREATE DATABASE IF NOT EXISTS report;

GRANT SELECT, DELETE, CREATE ON report.* TO 'report_generator_user'@'%';

USE report;

CREATE TABLE IF NOT EXISTS `report` (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    daily_payment_report INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);