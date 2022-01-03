package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

const (
	TREND_BULL    = "bull"
	TREND_UNKNOWN = "unknown"
	TREND_BEAR    = "bear"
)

type ReportService struct {
	application application.Application
}

func NewReportService(app application.Application) *ReportService {
	return &ReportService{
		application: app,
	}
}

func (r ReportService) GenerateReportStock(stock domain.Stock, quotes []domain.Quote) (*domain.ReportStock, error) {
	logger := r.application.Logger()
	periods := r.application.Periods()

	if len(quotes) == 0 {
		return nil, errors.New("No quotes found")
	}

	report := domain.ReportStock{
		Stock:     stock,
		BaseValue: quotes[0].Close,
		Periods:   make([]domain.Period, 0),
	}

	if len(quotes) == 0 {
		logger.Warn("No quotes fond", zap.String("stock", stock.Symbol))
		return &report, errors.New("no quotes found")
	}

	for i := range periods {
		movingAgerage, err := r.calculateMovingAveragePeriod(stock.Symbol, quotes, periods[i])
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		report.Periods = append(report.Periods, domain.Period{
			Name:  periods[i],
			Value: movingAgerage,
		})
	}

	return &report, nil
}

func (r ReportService) SaveReport(report domain.Report) error {
	// preparing header
	fileContent := "symbol,trend,"
	for i := range r.application.Periods() {
		fileContent += fmt.Sprintf("MA Period%d,", r.application.Periods()[i])
	}
	fileContent = fileContent[:len(fileContent)-1] + "\n"

	// preparing content
	for i := range report.Stocks {
		reportStock := report.Stocks[i]

		if reportStock.IsEmpty() {
			continue
		}

		fileContent += fmt.Sprintf("%s,", reportStock.Stock.Symbol)
		fileContent += fmt.Sprintf("%s,", r.identifyTrend(reportStock))

		for i := range reportStock.Periods {
			fileContent += fmt.Sprintf("%s,", reportStock.Periods[i].Value.StringFixed(2))
		}
		fileContent = fileContent[:len(fileContent)-1] + "\n"
	}

	// writing file
	now := time.Now()
	fileFullPath := fmt.Sprintf("%s/report_%s.csv", r.application.ReportPath(), now.Format("20060102150405"))
	ioutil.WriteFile(fileFullPath, []byte(fileContent), 0644)
	return nil
}

func (r ReportService) calculateMovingAveragePeriod(s string, quotes []domain.Quote, period int) (decimal.Decimal, error) {
	// calculating the simple moving average for the selected period
	result := decimal.NewFromInt(0)
	period = int(math.Min(float64(period), float64(len(quotes)))) // :-(

	for i := 0; i < period; i++ {
		result = result.Add(quotes[i].Close)
	}

	return result.Div(decimal.NewFromInt(int64(period))), nil
}

func (r ReportService) identifyTrend(reportStock domain.ReportStock) string {
	trend := 0

	for i := range reportStock.Periods {
		if i == 0 {
			continue
		}

		period := reportStock.Periods[i]
		prevPreriod := reportStock.Periods[i-1]

		if period.Value.Equal(decimal.Decimal{}) {
			break
		}

		if prevPreriod.Value.GreaterThanOrEqual(period.Value) {
			trend++
		} else {
			trend--
		}
	}

	if trend == len(reportStock.Periods)-1 {
		return TREND_BULL
	}

	if trend == (len(reportStock.Periods)-1)*-1 {
		return TREND_BEAR
	}

	return TREND_UNKNOWN
}
