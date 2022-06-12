package main

import (
	config "team3-task/config"
	"team3-task/internal/app"
	log "team3-task/pkg/logging"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Config error: %s", err)

	}

	app.Run(cfg)
}
