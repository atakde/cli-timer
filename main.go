package main

import (
	"fmt"
	"os"
	"time"
)

type CountDown struct {
	total int
	hour  int
	min   int
	sec   int
}

func getRemainingTime(endTime time.Time) CountDown {
	total := int(time.Until(endTime).Seconds())
	hour := int(time.Until(endTime).Hours())
	min := int(time.Until(endTime).Minutes())
	sec := int(time.Until(endTime).Seconds())

	return CountDown{total, hour, min, sec}
}

func Timer(duration time.Duration) {
	startTime := time.Now()
	endTime := startTime.Add(duration)

	// current time
	fmt.Println("Current time is", startTime.Format("15:04:05"))
	fmt.Println("Timer set for", duration)
	fmt.Println("Timer will end at", endTime.Format("15:04:05"))

	// ticker
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		// remaining time
		remainingTime := getRemainingTime(endTime)
		secForDisplay := remainingTime.sec % 60
		minForDisplay := remainingTime.min % 60
		hourForDisplay := remainingTime.hour % 24
		fmt.Printf("\rRemaining time: %02d:%02d:%02d", hourForDisplay, minForDisplay, secForDisplay)
		if time.Now().After(endTime) || time.Now().Equal(endTime) {
			// current time
			fmt.Println("\nCurrent time is", time.Now().Format("15:04:05"))
			ticker.Stop()
			break
		}
	}
}

func main() {
	timeArg := os.Args[1]
	duration, err := time.ParseDuration(timeArg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Timer(duration)
	os.Exit(0)
}
