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
	application   application.Application
	reportService *services.ReportService
}

func NewShinobi(application application.Application) Shinobi {
	return Shinobi{
		application:   application,
		reportService: services.NewReportService(application),
	}
}

func (s Shinobi) Sync() {
	logger := s.application.Logger()
	stocks := s.application.Stocks()
	report := domain.Report{}

	logger.Info(fmt.Sprint("Synching ", len(stocks), " stocks"))

	var wg sync.WaitGroup
	wg.Add(len(stocks))

	for i := range stocks {
		go func(i int) {
			stock := stocks[i]
			logger.Info("Synching stock " + stock.Symbol)
			if reportStock := s.syncStock(stock); reportStock != nil {
				report.AddStock(*reportStock)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	logger.Info("Stocks synchonized")

	if err := s.reportService.SaveReport(report); err != nil {
		logger.Error("Impossible to generate report", zap.Error(err))
	}
}

func (s Shinobi) syncStock(stock domain.Stock) *domain.ReportStock {
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

	report, err := s.reportService.GenerateReportStock(stock, quotes)
	if err != nil {
		logger.Error("Impossible to generate report", zap.String("symbol", stock.Symbol), zap.Error(err))
	}

	return report
}
