package app

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	config "team3-task/config"
	"team3-task/pkg/httpserver"
	log "team3-task/pkg/logging"
)

func Run(cfg *config.Config) {
	var err error
	// db
	strurl := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.PG.Username),
		url.QueryEscape(cfg.PG.Password),
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.DBName,
		cfg.PG.ConnTimeout)
	_ = strurl
	// insPgDB, err := pg.NewInsPgDB(strurl, cfg.PG.PoolMax)
	// if err != nil {
	// 	log.Fatal("Can't create DB connection: %v", err)
	// }
	// defer insPgDB.Close()

	//_ = insPgDB

	// HTTP Server
	//handler := .New()
	mux := http.NewServeMux()
	//v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(mux, cfg.HTTP)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error("app - Run - httpServer.Notify: %w", err)
		// case err = <-rmqServer.Notify():									// kafka
		// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown: %w", err)
	}

	// err = rmqServer.Shutdown()										// kafka
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	// }
}
