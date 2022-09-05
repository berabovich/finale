package support

import (
	"encoding/json"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// SupportGet read from link support and check data
func SupportGet() []SupportData {
	var supportData []SupportData
	r, err := http.Get("http://127.0.0.1:8383/support")

	if err != nil {
		return supportData
	}
	if r.StatusCode != http.StatusOK {
		return supportData
	}

	f, err := io.ReadAll(r.Body)
	err = json.Unmarshal(f, &supportData)

	return supportData
}
