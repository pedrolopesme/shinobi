package services

import (
	"math"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/ports"
	"github.com/pedrolopesme/shinobi/internal/utils"
	"go.uber.org/zap"
)

type AlphaVantageQuoteService struct {
	application domain.Application
	repositoy   ports.QuotesRepositoryContract
}

func NewAlphaVantageQuoteService(application domain.Application, repo ports.QuotesRepositoryContract) AlphaVantageQuoteService {
	return AlphaVantageQuoteService{
		application: application,
		repositoy:   repo,
	}
}

// GetQuotes returns the quotes from a given symbol
func (s AlphaVantageQuoteService) GetQuotes(symbol string) ([]domain.Quote, error) {
	// retrieving logger from application context
	logger := s.application.Logger()

	// trying to retrieve symbol data points
	quotes, err := s.repositoy.GetQuotes(symbol)
	if err != nil {
		logger.Error("impossible to retrieve quotes", zap.String("symbol", symbol), zap.Error(err))
		return nil, err
	}

	return quotes, nil
}

// TODO add some cache to avoid call the same symbol multiple times
func (a AlphaVantageQuoteService) GetMovingAveragePeriod(symbol string, period int) (float32, error) {
	logger := a.application.Logger()
	quotes, err := a.GetQuotes(symbol)
	if err != nil {
		logger.Error("impossible to calculate moving average", zap.String("symbol", symbol), zap.Error(err))
		return 0, err
	}

	// calculating the simple moving average for the selected period
	result := float32(0)
	period = int(math.Min(float64(period), float64(len(quotes)))) // :-(

	for i := 0; i < period; i++ {
		result += quotes[i].Close
	}

	result /= float32(period)
	return utils.RoundMoney(result)
}
