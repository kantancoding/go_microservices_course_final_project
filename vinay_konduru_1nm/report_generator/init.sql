CREATE USER 'payment_report_user'@'%' IDENTIFIED BY 'Auth123';

CREATE DATABASE payment_report;

GRANT INSERT, SELECT, UPDATE, DELETE ON payment_report.* TO 'payment_report_user'@'%';

USE payment_report;

CREATE TABLE `payment_report` (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  amount INT NOT NULL,
  report_date VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

