package main

import (
	"time"
	"log"
	"os"
	"os/signal"
	"github.com/robfig/cron/v3"
	reportGenerator "github.com/vinay-winai/go_microservices_course_final_project/internal/implementation"
)

func main() {
	defer log.Println("Service crashed")
	cronJob := cron.New()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	dailyPaymentsReport:= func() {
		log.Println("Running job at", time.Now())
		log.Println("Processing daily report...")
		res,err := reportGenerator.DailyPaymentsReport()
		log.Println(res)
		if err != nil {
			log.Println("Job failed:", err)
			quit <- os.Interrupt
			}
		}
	cron_id, err := cronJob.AddFunc("@midnight", dailyPaymentsReport)
	if err != nil {
		log.Println("Error scheduling job:", err)
		return
	}
	log.Printf("Scheduled job with ID %d\n", cron_id)
	cronJob.Start()
	<-quit
	jobCtx := cronJob.Stop()
	log.Println("Waiting for jobs to be done...")
	<-jobCtx.Done()
	err = jobCtx.Err()
	if err != nil {
		log.Println("Cause of failure:", err)
	}
}