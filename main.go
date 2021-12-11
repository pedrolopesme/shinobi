package main

import (
	"github.com/pedrolopesme/shinobi/cmd"
	"github.com/pedrolopesme/shinobi/internal/domain"
	"go.uber.org/zap"
)

func main() {
	application := domain.Application{
		AlphaVantageAPIKey: "VIV654H7KZ7VHL5V",
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	shinobi := cmd.NewShinobi(application, *logger)
	shinobi.Run()
}
