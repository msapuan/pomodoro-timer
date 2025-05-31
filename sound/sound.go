package sound

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlaySound(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("Gagal membuka file suara:", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Println("Gagal decode file MP3:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
