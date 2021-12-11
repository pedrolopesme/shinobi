package ports

import "github.com/pedrolopesme/shinobi/internal/domain"

type QuoteContract interface {
	GetQuotes(symbol string) ([]domain.Quote, error)
	GetMovingAveragePeriod(symbol string, period int) (float32, error)
}
