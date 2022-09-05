package incident

import (
	"encoding/json"
	"io"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}

// IncidentGet get incident data from link and check it
func IncidentGet() []IncidentData {
	var incidentData []IncidentData
	r, err := http.Get("http://127.0.0.1:8383/accendent")

	if err != nil {
		return incidentData
	}
	if r.StatusCode != http.StatusOK {
		return incidentData
	}

	f, err := io.ReadAll(r.Body)
	err = json.Unmarshal(f, &incidentData)

	return incidentData
}
