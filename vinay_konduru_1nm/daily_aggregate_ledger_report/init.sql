CREATE USER 'ledger_report_user'@'%' IDENTIFIED BY 'Auth123';

CREATE DATABASE daily_aggregate_ledger_report;

GRANT INSERT, SELECT, UPDATE, DELETE ON daily_aggregate_ledger_report.* TO 'ledger_report_user'@'%';

USE daily_aggregate_ledger_report;

CREATE TABLE `daily_aggregate_ledger_report` (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  daily_aggregate INT NOT NULL,
  aggregated_date VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

