package services

import (
	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"github.com/pedrolopesme/shinobi/internal/ports"
	"go.uber.org/zap"
)

type AlphaVantageQuoteService struct {
	application application.Application
	repositoy   ports.QuotesRepositoryContract
}

func NewAlphaVantageQuoteService(application application.Application, repo ports.QuotesRepositoryContract) AlphaVantageQuoteService {
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
