package services

import (
	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/ports"
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

// GetQuotes returns the last 30 quotes from a given symbol
func (s AlphaVantageQuoteService) GetQuotes(symbol string) ([]domain.Quote, error) {
	// retrieving logger from application context
	logger := s.application.Logger()

	// trying to retrieve symbol data points
	quotes, err := s.repositoy.GetQuotes(symbol)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// returning last 30 quotes
	if len(quotes) > 0 {
		return quotes[:30], nil
	}

	return nil, nil
}
