package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func testPrint() {
	fmt.Println("test Cron ", time.Now())
}

func TestGocron() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(testPrint)

	s.Every("5m").Do(testPrint)

	s.Every(5).Days().Do(testPrint)

	s.Every(1).Month(1, 2, 3).Do(testPrint)

	s.Every(1).Day().At("10:30").Do(testPrint)

	s.Every(1).Day().At("10:30;08:00").Do(testPrint)

	s.Every(1).Day().At("10:30").At("08:00").Do(testPrint)

	s.Every(1).MonthLastDay().Do(testPrint)

	s.Every(2).MonthLastDay().Do(testPrint)

	s.Cron("*/1 * * * *").Do(testPrint)

	s.StartAsync()

	s.StartBlocking()
}
