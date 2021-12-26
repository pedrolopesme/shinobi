package application

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"go.uber.org/zap"
)

const (
	CTX_APP          = "CTX_Application"
	CTX_ALPHAVANTAGE = "CTX_ALPHAVANTAGE"
	CTX_LOGGER       = "CTX_LOGGER"
)

type Application struct {
	ctx    context.Context
	config Config
}

func NewApplication(
	configPath string,
	logger zap.Logger) *Application {

	appConfig, err := loadConfigFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, CTX_ALPHAVANTAGE, appConfig.AlphaVantage.ApiKey)
	ctx = context.WithValue(ctx, CTX_LOGGER, logger)

	return &Application{
		ctx:    ctx,
		config: *appConfig,
	}
}

func loadConfigFile(path string) (*Config, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var appConfig Config
	if err := json.Unmarshal(fileContent, &appConfig); err != nil {
		return nil, err
	}
	return &appConfig, nil
}

func (a Application) Logger() zap.Logger {
	return a.ctx.Value(CTX_LOGGER).(zap.Logger)
}

func (a Application) AlphaVantageKey() string {
	return a.ctx.Value(CTX_ALPHAVANTAGE).(string)
}

func (a Application) Stocks() []domain.Stock {
	return a.config.Stocks
}

func (a Application) Periods() []int {
	return a.config.Report.Periods
}

func (a Application) ReportPath() string {
	return a.config.Report.Path
}
