package main

import (
	"time"
	"log"
	lr "github.com/vinay-winai/go_microservices_course_final_project/internal/implementation"
)

func main() {
	defer log.Println("Service crashed")
	for {
		time_to_midnight := time.Until(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local))
		log.Println("Sleeping for", time_to_midnight)
		time.Sleep(time_to_midnight)
		log.Println("Processing daily report...")
		res,err := lr.GenerateDailyAggregateLedgerReport()
		log.Println(res)
		if err != nil {
			log.Println(err)
		}
	}
}