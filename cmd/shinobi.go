package cmd

import (
	"fmt"
	"os"
	"sync"

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
	logger := s.application.Logger()
	stocks := s.application.Stocks()

	logger.Info(fmt.Sprint("Synching ", len(stocks), " stocks"))

	var wg sync.WaitGroup
	wg.Add(len(stocks))

	for i := range stocks {
		go func(i int) {
			stock := stocks[i]
			logger.Info("Synching stock " + stock.Symbol)
			s.syncStock(stock)
			wg.Done()
		}(i)
	}

	wg.Wait()
	logger.Info("Stocks synchonized")

}

func (s Shinobi) syncStock(stock domain.Stock) {
	logger := s.application.Logger()
	logger.Info("Running Shinobi on " + stock.Symbol)

	quotesRepo := quotes.NewAlphaVantageQuoteRepository(s.application)
	quotesService := services.NewAlphaVantageQuoteService(s.application, quotesRepo)
	reportService := services.NewReportService(s.application)

	logger.Info("Getting quotes")
	quotes, err := quotesService.GetQuotes(stock.Symbol)
	if err != nil {
		logger.Error("impossible to calculate moving average", zap.String("symbol", stock.Symbol), zap.Error(err))
		os.Exit(1)
	}

	report, err := reportService.GenerateReport(stock, quotes)
	if err != nil {
		logger.Error("Impossible to generate report", zap.String("symbol", stock.Symbol), zap.Error(err))
	}

	if err := reportService.SaveReport(*report); err != nil {
		logger.Error("Impossible to generate report", zap.String("symbol", stock.Symbol), zap.Error(err))
	}
}
