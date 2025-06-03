package sound

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Path absolut ke direktori tempat file suara disimpan oleh .deb
const systemSoundPath = "/usr/share/pomodoro-timer/sounds"

func PlaySound(filename string) {
	fullPath := filepath.Join(systemSoundPath, filename)

	f, err := os.Open(fullPath)
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
