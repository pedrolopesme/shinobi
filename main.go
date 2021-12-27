package main

import (
	"flag"

	"github.com/pedrolopesme/shinobi/cmd"
	"github.com/pedrolopesme/shinobi/internal/domain/application"
	"go.uber.org/zap"
)

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	application := application.NewApplication(
		*configFile,
		*logger,
	)

	cmd.NewShinobi(*application).Run()
}
