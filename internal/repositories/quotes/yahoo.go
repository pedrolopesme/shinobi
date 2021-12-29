package quotes

import (
	"time"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"go.uber.org/zap"
)

const (
	MAX_QUOTES_DAY = 200
)

type YahooQuotesRepository struct {
	application application.Application
}

func NewYahooQuotesRepository(application application.Application) *YahooQuotesRepository {
	return &YahooQuotesRepository{
		application: application,
	}
}

func (a YahooQuotesRepository) GetQuotes(symbol string) ([]domain.Quote, error) {
	logger := a.application.Logger()
	logger.Info("Retrieving Quotes from Yahoo Finance API", zap.String("symbol", symbol))

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Impossible to get quotes", zap.String("symbol", symbol))
		}
	}()

	now := time.Now()
	dateStart := now.Add(-1 * MAX_QUOTES_DAY * 24 * time.Hour)

	p := &chart.Params{
		Symbol:   symbol,
		Start:    datetime.New(&dateStart),
		End:      datetime.New(&now),
		Interval: datetime.OneDay,
	}

	quotes := make([]domain.Quote, 0)
	tickCursor := chart.Get(p)

	for tickCursor.Next() {
		tick := tickCursor.Bar()

		quote := domain.Quote{
			Date:   time.Unix(int64(tick.Timestamp), 0),
			Open:   tick.Open,
			High:   tick.High,
			Low:    tick.Low,
			Close:  tick.Close,
			Volume: int32(tick.Volume),
		}

		quotes = append(quotes, quote)
	}

	return reverseQuotes(quotes), nil
}

func reverseQuotes(quotes []domain.Quote) []domain.Quote {
	for i := 0; i < len(quotes)/2; i++ {
		j := len(quotes) - i - 1
		quotes[i], quotes[j] = quotes[j], quotes[i]
	}
	return quotes
}
