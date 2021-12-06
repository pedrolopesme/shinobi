package cmd

import (
	"context"

	"github.com/pedrolopesme/shinobi/internal/domain"
	"github.com/pedrolopesme/shinobi/internal/repositories/quotes"
	"github.com/pedrolopesme/shinobi/internal/services"
	"go.uber.org/zap"
)

type Shinobi struct {
	ctx context.Context
}

func NewShinobi(
	application domain.Application,
	logger zap.Logger,
) Shinobi {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.CTX_APP, application)
	ctx = context.WithValue(ctx, domain.CTX_LOGGER, logger)
	return Shinobi{ctx: ctx}
}

func (s Shinobi) Run() {
	logger := s.ctx.Value(domain.CTX_LOGGER).(zap.Logger)

	symbol := "MGLU3.SA"
	logger.Info("Running Shinobi on " + symbol)

	quotesRepo := quotes.NewAlphaVantageQuoteRepository(s.ctx)
	quotesService := services.NewAlphaVantageQuoteService(s.ctx, quotesRepo)

	logger.Info("Getting quotes")
	quotesService.GetQuotes(symbol)
}
