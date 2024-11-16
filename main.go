package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// helloWorldHandler is an HTTP handler that responds with "Hello, World!"
type helloWorldHandler struct{}

// ServeHTTP handles HTTP requests for helloWorldHandler
func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug().Str("method", r.Method).Str("path", r.URL.Path).Msg("Handled request")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		log.Error().Err(err).Msg("Failed to write response")
		return
	}
}

func main() {

	addr := flag.String("addr", ":8080", "address to listen on")
	readTimeout := flag.Duration("read-timeout", 10*time.Second, "read timeout")
	writeTimeout := flag.Duration("write-timeout", 10*time.Second, "write timeout")
	logLevel := flag.String("log-level", "info", "log level (debug, info, warn, error, fatal, panic)")
	flag.Parse()

	lvl, err := zerolog.ParseLevel(*logLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse log level")
	}
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(lvl)

	s := &http.Server{
		Addr:         *addr,
		Handler:      &helloWorldHandler{},
		ReadTimeout:  *readTimeout,
		WriteTimeout: *writeTimeout,
	}
	log.Info().Msg("Server started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	serverErr := make(chan error, 1)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	var shutdownOnce sync.Once
	exitCode := 0

	select {
	case err := <-serverErr:
		log.Error().Err(err).Msg("Server encountered an error")
		exitCode = 1
		shutdownOnce.Do(func() {
			shutdownServer(s)
		})
	case <-stop:
		log.Info().Msg("Shutting down server...")
		shutdownOnce.Do(func() {
			shutdownServer(s)
		})
	}

	log.Info().Msg("Server stopped")
	os.Exit(exitCode)
}

func shutdownServer(s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Warn().Msg("Server shutdown canceled")
		} else {
			log.Error().Err(err).Msg("Server shutdown failed")
		}
	} else {
		log.Info().Msg("Server shutdown complete")
	}
}
