package ports

import "github.com/pedrolopesme/shinobi/internal/domain"

type QuotesRepositoryContract interface {
	GetQuotes(symbol string) ([]domain.Quote, error)
}
