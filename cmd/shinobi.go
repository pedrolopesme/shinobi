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

	logger.Info("Getting quotes")
	quotes, err := quotesService.GetQuotes(stock.Symbol)
	if err != nil {
		logger.Error("impossible to calculate moving average", zap.String("symbol", stock.Symbol), zap.Error(err))
		os.Exit(1)
	}

	yesterday := quotes[0].Close
	lastWeek, _ := quotesService.GetMovingAveragePeriod(quotes, 5)
	lastMonth, _ := quotesService.GetMovingAveragePeriod(quotes, 20)
	lastQuarter, _ := quotesService.GetMovingAveragePeriod(quotes, 60)
	last200Days, _ := quotesService.GetMovingAveragePeriod(quotes, 200)

	fmt.Println(stock.Symbol, "yesterday", yesterday, "lastWeek", lastWeek, "lastMonth", lastMonth, "lastQuarter", lastQuarter, "last200Days", last200Days)
}
