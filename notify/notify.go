package notify

import (
	"fmt"
	"os/exec"
)

func Show(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Gagal mengirim notifikasi:", err)
	}
}