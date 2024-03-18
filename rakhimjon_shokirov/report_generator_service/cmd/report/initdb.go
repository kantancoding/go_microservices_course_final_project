package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbDriver = "mysql"

func initDB(user, password, host, port, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
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
		log.Fatalf("Error pinging to a db: %v", err)
	}
	log.Printf("pinged to a %s successfully", dbName)

	return db, nil
}
