package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Quote struct {
	Date   time.Time       `json:"date"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume int32           `json:"volume"`
}
