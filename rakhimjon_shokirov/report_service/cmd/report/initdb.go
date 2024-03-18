package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	dbDriver     = "mysql"
	dbReportName = "report"
	dbLedgerName = "ledger"
)

func initReportDB() (*sql.DB, error) {
	dbUser := "report_user"
	dbPassword := "Auth123"

	dsn := fmt.Sprintf("%s:%s@tcp(mysql-report:3306)/%s", dbUser, dbPassword, dbReportName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing db: %s", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging to the report db: %v", err)
	}
	log.Println("pinged to a report db successfully")

	return db, nil
}

func initLedgerReadOnlyDB() (*sql.DB, error) {
	dbUser := "ledgerReadOnlyModeUser"
	dbPassword := "ledgerReadOnlyMode123"

	dsn := fmt.Sprintf("%s:%s@tcp(mysql-ledger:3306)/%s", dbUser, dbPassword, dbLedgerName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing db: %s", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging to the ledger db: %v", err)
	}
	log.Println("pinged to a ledger db successfully")

	return db, nil
}
