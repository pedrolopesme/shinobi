package services

import (
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"github.com/pedrolopesme/shinobi/internal/utils"
)

type ReportService struct {
	application application.Application
}

func NewReportService(app application.Application) *ReportService {
	return &ReportService{
		application: app,
	}
}

func (r ReportService) GenerateReport(stock domain.Stock, quotes []domain.Quote) (*domain.Report, error) {
	logger := r.application.Logger()
	periods := r.application.Periods()

	report := domain.Report{
		Stock:   stock,
		Periods: make([]domain.Period, 0),
	}

	for i := range periods {
		movingAgerage, err := r.calculateMovingAveragePeriod(quotes, periods[i])
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
	fileContent := "symbol;"
	for i := range r.application.Periods() {
		fileContent += fmt.Sprintf("D%d;", r.application.Periods()[i])
	}
	fileContent = fileContent[:len(fileContent)-1] + "\n"

	// preparing content
	fileContent += report.Stock.Symbol
	for i := range report.Periods {
		fileContent += fmt.Sprintf("%f;", report.Periods[i].Value)
	}
	fileContent = fileContent[:len(fileContent)-1] + "\n"

	// writing file
	now := time.Now()
	fileFullPath := fmt.Sprintf("%s/report_%s.csv", r.application.ReportPath(), now.Format("20060102150405"))
	ioutil.WriteFile(fileFullPath, []byte(fileContent), 0644)
	return nil
}

func (r ReportService) calculateMovingAveragePeriod(quotes []domain.Quote, period int) (float32, error) {
	// calculating the simple moving average for the selected period
	result := float32(0)
	period = int(math.Min(float64(period), float64(len(quotes)))) // :-(

	for i := 0; i < period; i++ {
		result += quotes[i].Close
	}

	result /= float32(period)
	return utils.RoundMoney(result)
}
