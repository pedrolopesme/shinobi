package main

import (
	"github.com/pedrolopesme/shinobi/cmd"
	"github.com/pedrolopesme/shinobi/internal/domain"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	application := domain.NewApplication(
		"VIV654H7KZ7VHL5V",
		*logger,
	)

	cmd.NewShinobi(*application).Run()
}
