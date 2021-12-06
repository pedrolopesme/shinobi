package services

import (
	"context"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/ports"
	"go.uber.org/zap"
)

type AlphaVantageQuoteService struct {
	ctx       context.Context
	repositoy ports.QuotesRepositoryContract
}

func NewAlphaVantageQuoteService(ctx context.Context, repo ports.QuotesRepositoryContract) AlphaVantageQuoteService {
	return AlphaVantageQuoteService{
		ctx:       ctx,
		repositoy: repo,
	}
}

// GetQuotes returns the last 30 quotes from a given symbol
func (s AlphaVantageQuoteService) GetQuotes(symbol string) ([]domain.Quote, error) {
	// retrieving logger from application context
	logger := s.ctx.Value(domain.CTX_LOGGER).(zap.Logger)

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
