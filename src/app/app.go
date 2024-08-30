package app

import (
	"fmt"
	"github.com/rs/zerolog"
	"go-template/src/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.MessageFieldName = "msg"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	mainApi := api.New()

	server := http.Server{
		Addr:    ":8080",
		Handler: mainApi.Handler(),
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for api requests.
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	}
}
