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
		"INSERT YOU KEY HERE",
		*logger,
	)

	cmd.NewShinobi(*application).Run()
}
