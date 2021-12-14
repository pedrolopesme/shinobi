package application

import "github.com/pedrolopesme/shinobi/internal/domain"

type Config struct {
	AlphaVantage AlphaVantage   `json:"alphavantage"`
	Database     Database       `json:"database"`
	Stocks       []domain.Stock `json:"stocks"`
}
