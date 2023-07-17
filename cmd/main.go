package main

import (
	"flag"
	"log"

	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/server"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
)

// @contact.name Ashkan Maleki
// @contact.url https://github.com/ashkan-maleki/go-es-cqrs
// @contact.email no email
func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName(server.GetMicroserviceName(cfg))
	appLogger.Fatal(server.NewServer(cfg, appLogger).Run())
}
