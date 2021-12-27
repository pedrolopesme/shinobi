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

const (
	MAX_CONCURRENT_WORKERS = 1
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

func (s Shinobi) Run() {
	logger := s.application.Logger()
	stocks := s.application.Stocks()
	numJobs := len(stocks)

	report := domain.Report{}

	logger.Info(fmt.Sprint("Synching ", len(stocks), " stocks"))
	stocksQueue := make(chan domain.Stock, numJobs)
	results := make(chan domain.ReportStock, numJobs)

	for i := 0; i < MAX_CONCURRENT_WORKERS; i++ {
		go s.syncWorker(stocksQueue, results)
	}

	for i := 0; i < numJobs; i++ {
		stocksQueue <- stocks[i]
	}
	close(stocksQueue)

	for i := 0; i < numJobs; i++ {
		report.AddStock(<-results)
	}
	logger.Info("Stocks synchonized")

	s.reportService.SaveReport(report)
	logger.Info("Report saved")
}

func (s Shinobi) syncWorker(stocks <-chan domain.Stock, results chan<- domain.ReportStock) {
	logger := s.application.Logger()
	for stock := range stocks {
		logger.Info("Synching stock " + stock.Symbol)
		if reportStock := s.syncStock(stock); reportStock != nil {
			results <- *reportStock
		}
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
