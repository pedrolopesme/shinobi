package cmd

import (
	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/repositories/quotes"
	"github.com/pedrolopesme/shinobi/internal/services"
)

type Shinobi struct {
	application domain.Application
}

func NewShinobi(application domain.Application) Shinobi {
	return Shinobi{application: application}
}

func (s Shinobi) Run() {
	logger := s.application.Logger()

	symbol := "MGLU3.SA"
	logger.Info("Running Shinobi on " + symbol)

	quotesRepo := quotes.NewAlphaVantageQuoteRepository(s.application)
	quotesService := services.NewAlphaVantageQuoteService(s.application, quotesRepo)

	logger.Info("Getting quotes")
	quotesService.GetQuotes(symbol)
}
