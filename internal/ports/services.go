package ports

import "github.com/pedrolopesme/shinobi/internal/domain"

type QuoteContract interface {
	GetQuotes(symbol string) ([]domain.Quote, error)
}

type ReportServiceContract interface {
	GenerateReport(stock domain.Stock, quotes []domain.Quote) (*domain.Report, error)
	SaveReport(report domain.Report) error
}
