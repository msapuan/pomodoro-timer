package timer

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

func StartTimer(minutes int, label string) {
	duration := minutes * 60 // in seconds
	fmt.Printf("\n⏳ %s selama %d menit\n", label, minutes)

	bar := progressbar.NewOptions(duration,
		progressbar.OptionSetDescription(label),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "#", SaucerPadding: "-", BarStart: "[", BarEnd: "]"}),
	)

	for i := 0; i < duration; i++ {
		bar.Add(1)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n✅ %s selesai!\n", label)
}
