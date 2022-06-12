package httpserver

import (
	"context"
	"net/http"
	"team3-task/config"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":10000"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, cfg config.HTTP) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	if time.Duration(cfg.HttpReadTimeout) != _defaultReadTimeout {
		httpServer.ReadTimeout = time.Duration(cfg.HttpReadTimeout)
	}
	if time.Duration(cfg.HttpWriteTimeout) != _defaultWriteTimeout {
		httpServer.WriteTimeout = time.Duration(cfg.HttpWriteTimeout)
	}
	if cfg.HttpAddr != _defaultAddr {
		httpServer.Addr = cfg.HttpAddr
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
