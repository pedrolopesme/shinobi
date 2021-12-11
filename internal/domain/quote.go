package domain

import "time"

type Quote struct {
	Date   time.Time `json:"date"`
	Open   float32   `json:"open"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
	Close  float32   `json:"close"`
	Volume int32     `json:"volume"`
}
