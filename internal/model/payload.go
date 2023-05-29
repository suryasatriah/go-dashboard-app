package model

type Payload struct {
	Voltage string `json:"voltage"`
	Power   string `json:"power"`
	Status  bool   `json:"status"`
}
