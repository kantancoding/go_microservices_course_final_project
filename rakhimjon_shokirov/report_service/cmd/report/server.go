package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	"go_microservices_course_final_project/rakhimjon_shokirov/report_service/internal/implementation"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
)

var (
	reportDb, ledgerDb *sql.DB
)

type application struct {
	reportDb *sql.DB
	ledgerDb *sql.DB
}

func main() {
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)
	app := application{}

	// Open a database connection
	rDB, err := initReportDB()
	if err != nil {
		log.Fatal("error initilizing rDB: ", err.Error())
	}
	app.reportDb = rDB

	lDb, err := initLedgerReadOnlyDB()
	if err != nil {
		log.Fatal("error initilizing lDb: ", err.Error())
	}
	app.ledgerDb = lDb

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		OSCall := <-quitSignal
		log.Printf("\nSystem Call: %+v", OSCall)
		cancel()
	}()

	app.run(ctx)
}

func (app *application) run(ctx context.Context) {
	c := cron.New()
	if err := c.AddFunc("0 0 10 * * *", func() {
		log.Println("Executing task at 10:00 AM", time.Now())

		dailyAmount, err := implementation.GetDailyAmount(ctx, app.ledgerDb)
		if err != nil {
			log.Printf("error on implementation.GetDailyAmount %v", err.Error())
			c.Stop()
			return
		}

		if err := implementation.Insert(ctx, app.reportDb, dailyAmount); err != nil {
			log.Printf("error on report.Insert %v", err.Error())
			c.Stop()
			return
		}
	}); err != nil {
		log.Fatalf("error on c.AddFunc: %+v", err.Error())
	}

	go c.Start()

	<-ctx.Done()
	c.Stop()
	log.Println("cron job is shut down")
}
