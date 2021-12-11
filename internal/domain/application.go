package domain

import (
	"context"

	"go.uber.org/zap"
)

const (
	CTX_APP          = "CTX_Application"
	CTX_ALPHAVANTAGE = "CTX_ALPHAVANTAGE"
	CTX_LOGGER       = "CTX_LOGGER"
)

type Application struct {
	ctx context.Context
}

func NewApplication(
	alphaVantageAPIKey string,
	logger zap.Logger) *Application {

	ctx := context.Background()
	ctx = context.WithValue(ctx, CTX_ALPHAVANTAGE, alphaVantageAPIKey)
	ctx = context.WithValue(ctx, CTX_LOGGER, logger)

	return &Application{
		ctx: ctx,
	}
}

func (a Application) Logger() zap.Logger {
	return a.ctx.Value(CTX_LOGGER).(zap.Logger)
}

func (a Application) AlphaVantageKey() string {
	return a.ctx.Value(CTX_ALPHAVANTAGE).(string)
}
