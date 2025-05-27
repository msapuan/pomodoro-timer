package main

import (
	"flag"
	"fmt"
	"pomodoro-timer/quotes"
	"pomodoro-timer/timer"
)

func main() {
	work := flag.Int("work", 25, "Durasi waktu kerja (menit)")
	breakTime := flag.Int("break", 5, "Durasi istirahat (menit)")
	flag.Parse()

	fmt.Println("🌟 Motivasi:")
	fmt.Println(quotes.GetQuote())

	timer.StartTimer(*work, "Waktu Fokus")
	timer.StartTimer(*breakTime, "Waktu Istirahat")
}
