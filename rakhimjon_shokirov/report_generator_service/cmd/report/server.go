package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"report_generator_service/config"
	"report_generator_service/internal/implementation"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"github.com/sethvargo/go-envconfig"
)

type application struct {
	reportDb *sql.DB
	ledgerDb *sql.DB
}

func main() {
	var (
		quitSignal = make(chan os.Signal, 1)
		app        application
		cfg        config.Config
	)

	signal.Notify(quitSignal, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		OSCall := <-quitSignal
		log.Printf("\nSystem Call: %+v", OSCall)
		cancel()
	}()

	if err := envconfig.Process(context.TODO(), &cfg); err != nil {
		log.Fatal(err)
	}

	// Open a report database connection
	rDB, err := initDB(
		cfg.ReportMysql.User,
		cfg.ReportMysql.Password,
		cfg.ReportMysql.Host,
		cfg.ReportMysql.Port,
		cfg.ReportMysql.Database,
	)
	if err != nil {
		log.Fatal("error initilizing rDB: ", err.Error())
	}
	app.reportDb = rDB

	// Open a ledger database connection
	lDb, err := initDB(
		cfg.LedgerMysql.User,
		cfg.LedgerMysql.Password,
		cfg.LedgerMysql.Host,
		cfg.LedgerMysql.Port,
		cfg.LedgerMysql.Database,
	)
	if err != nil {
		log.Fatal("error initilizing lDb: ", err.Error())
	}
	app.ledgerDb = lDb

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
