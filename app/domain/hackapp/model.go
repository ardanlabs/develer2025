package hackapp

import "encoding/json"

type Status struct {
	Status string
}

func (app Status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}
