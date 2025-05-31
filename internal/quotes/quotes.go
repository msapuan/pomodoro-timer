package quotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func GetQuote() string {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		return "Tetap semangat! Kamu pasti bisa!"
	}
	defer resp.Body.Close()

	var result []Quote
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result) == 0 {
		return "Terus maju, jangan menyerah!"
	}
	return fmt.Sprintf("\"%s\" - %s", result[0].Q, result[0].A)
}