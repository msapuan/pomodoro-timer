package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"pomodoro-timer/notify"
	"pomodoro-timer/quotes"
	"pomodoro-timer/timer"
)

func getInputInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if num, err := strconv.Atoi(text); err == nil && num > 0 {
			return num
		}
		fmt.Println("⚠️  Masukkan angka yang valid.")
	}
}

func main() {
	// Input langsung dari user satu kali
	work := getInputInt("⏱️  Waktu fokus (menit): ")
	breakTime := getInputInt("🛌 Waktu istirahat (menit): ")
	cycles := getInputInt("🔁 Jumlah siklus Pomodoro: ")

	fmt.Println("\n🌟 Motivasi:")
	fmt.Println(quotes.GetQuote())

	for i := 1; i <= cycles; i++ {
		fmt.Printf("\n🔄 Siklus %d dari %d\n", i, cycles)

		timer.StartTimer(work, "Waktu Fokus")
		notify.Show("Pomodoro Timer", "Sesi fokus selesai. Saatnya istirahat!")

		// Hanya lakukan istirahat jika belum di siklus terakhir
		if i < cycles {
			timer.StartTimer(breakTime, "Waktu Istirahat")
			notify.Show("Pomodoro Timer", "Istirahat selesai. Ayo mulai lagi!")
		}
	}

	fmt.Println("✅ Semua siklus selesai! Kerja bagus!")
}
