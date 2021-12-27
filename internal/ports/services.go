package ports

import "github.com/pedrolopesme/shinobi/internal/domain"

type QuoteContract interface {
	GetQuotes(symbol string) ([]domain.Quote, error)
}

type ReportServiceContract interface {
	GenerateReportStock(stock domain.Stock, quotes []domain.Quote) (*domain.ReportStock, error)
	SaveReport(report domain.Report) error
}
