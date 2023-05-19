package http

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	router http.Handler
	logger *zap.SugaredLogger
}

func New(router http.Handler, logger *zap.SugaredLogger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

func (s Server) Run() {
	// Startup with graceful shutdown
	srv := &http.Server{
		Addr:         "localhost:5000",
		Handler:      s.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		// Listen for SIGNINT and SIGTERM signals
		// and write them in quit channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		s.logger.Infoln("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)

		if err != nil {
			shutdownError <- err
		}

		s.logger.Infoln("Completing background tasks...")

		shutdownError <- nil
	}()

	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		s.logger.Fatal(err)
	}

	err = <-shutdownError
	if err != nil {
		s.logger.Fatal(err)
	}

	s.logger.Infoln("Stopped server")
}
