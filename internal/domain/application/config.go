package application

import "github.com/pedrolopesme/shinobi/internal/domain"

type ConfigReport struct {
	Path    string `json:"path"`
	Periods []int  `json:"periods"`
}

type Config struct {
	Stocks []domain.Stock `json:"stocks"`
	Report ConfigReport   `json:"report"`
}
