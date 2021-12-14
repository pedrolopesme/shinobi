package cmd

import (
	"fmt"
	"os"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"github.com/pedrolopesme/shinobi/internal/repositories/quotes"
	"github.com/pedrolopesme/shinobi/internal/services"
	"go.uber.org/zap"
)

type Shinobi struct {
	application application.Application
}

func NewShinobi(application application.Application) Shinobi {
	return Shinobi{application: application}
}

func (s Shinobi) Sync() {
	stocks := s.application.Stocks()
	for i := range stocks {
		s.syncStock(stocks[i])
	}
}

func (s Shinobi) syncStock(stock domain.Stock) {
	logger := s.application.Logger()
	logger.Info("Running Shinobi on " + stock.Symbol)

	quotesRepo := quotes.NewAlphaVantageQuoteRepository(s.application)
	quotesService := services.NewAlphaVantageQuoteService(s.application, quotesRepo)

	logger.Info("Getting quotes")
	quotes, err := quotesService.GetQuotes(stock.Symbol)
	if err != nil {
		logger.Error("impossible to calculate moving average", zap.String("symbol", stock.Symbol), zap.Error(err))
		os.Exit(1)
	}

	today := 0 // TODO today
	yesterday, _ := quotesService.GetMovingAveragePeriod(quotes, 1)
	lastWeek, _ := quotesService.GetMovingAveragePeriod(quotes, 5)
	lastMonth, _ := quotesService.GetMovingAveragePeriod(quotes, 20)
	lastQuarter, _ := quotesService.GetMovingAveragePeriod(quotes, 60)
	last200Days, _ := quotesService.GetMovingAveragePeriod(quotes, 200)

	fmt.Println(stock.Symbol, "today ", today, "yesterday", yesterday, "lastWeek", lastWeek, "lastMonth", lastMonth, "lastQuarter", lastQuarter, "last200Days", last200Days)
}
