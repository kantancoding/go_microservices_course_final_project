package daily_aggregate_ledger_report

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	dbDriver                    = "mysql"
	ledgerReportDbName          = "daily_aggregate_ledger_report"
	ledgerDbName                = "ledger"
	insertReportQuery           = "INSERT INTO daily_aggregate_ledger_report (daily_aggregate, aggregated_date) VALUES (?, ?)"
	dailyAggregateQuery         = "SELECT SUM(amount) AS daily_aggregate FROM ledger WHERE DATE(transaction_date) = ?"
	ledgerDbConnSubString       = "ledger"
	ledgerReportDbConnSubString = "daily-aggregate-ledger-report"
)

var db *sql.DB

func dbInit(dbName string, connSubString string) (db *sql.DB) {
	var err error

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	// Open a connection to the database
	dsn := fmt.Sprintf("%s:%s@tcp(mysql-%s:3306)/%s", dbUser, dbPassword, connSubString, dbName)
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	// check if the database is alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GenerateDailyAggregateLedgerReport() (string, error) {
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing ledger database: %s", err)
		}
	}()
	db = dbInit(ledgerDbName, ledgerDbConnSubString)
	var (
		yesterday      = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		dailyAggregate int64
		failure        = fmt.Sprintf("Failed to generate report for %s", yesterday)
	)
	err := db.QueryRow(dailyAggregateQuery, yesterday).Scan(&dailyAggregate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return failure, sql.ErrNoRows
		}
		return failure, err
	}
	err = storeLedgerReport(dailyAggregate, yesterday)
	if err != nil {
		return failure, err
	}
	return fmt.Sprintf("Daily aggregate for %s is %d", yesterday, dailyAggregate), nil
}

func storeLedgerReport(dailyAggregate int64, yesterday string) error {
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing daily_aggregate_ledger_report database: %s", err)
		}
	}()
	db = dbInit(ledgerReportDbName, ledgerReportDbConnSubString)
	stmt, err := db.Prepare(insertReportQuery)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dailyAggregate, yesterday)
	log.Printf("result of insert query execution in daily_aggregate_ledger_report_db is:%v", res)
	return err
}
