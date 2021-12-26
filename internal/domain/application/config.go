package application

import "github.com/pedrolopesme/shinobi/internal/domain"

type ConfigAlphaVantage struct {
	ApiKey string `json:"api_key"`
}
type ConfigReport struct {
	Path    string `json:"path"`
	Periods []int  `json:"periods"`
}

type Config struct {
	AlphaVantage ConfigAlphaVantage `json:"alphavantage"`
	Stocks       []domain.Stock     `json:"stocks"`
	Report       ConfigReport       `json:"report"`
}
