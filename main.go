package main

import (
	"flag"
	"fmt"
	"pomodoro-timer/notify"
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
	notify.Show("Pomodoro", "Sesi fokus selesai. Saatnya istirahat!")

	timer.StartTimer(*breakTime, "Waktu Istirahat")
	notify.Show("Pomodoro", "Istirahat selesai. Ayo mulai lagi!")

}
