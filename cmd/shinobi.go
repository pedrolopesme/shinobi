package cmd

import (
	"fmt"

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
	today := 0 // TODO today
	yesterday, _ := quotesService.GetMovingAveragePeriod(symbol, 1)
	lastWeek, _ := quotesService.GetMovingAveragePeriod(symbol, 5)
	lastMonth, _ := quotesService.GetMovingAveragePeriod(symbol, 20)
	lastQuarter, _ := quotesService.GetMovingAveragePeriod(symbol, 60)
	last200Days, _ := quotesService.GetMovingAveragePeriod(symbol, 200)
	fmt.Println(symbol, "today ", today, "yesterday", yesterday, "lastWeek", lastWeek, "lastMonth", lastMonth, "lastQuarter", lastQuarter, "last200Days", last200Days)
}
