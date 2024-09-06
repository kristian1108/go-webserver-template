package app

import (
	"context"
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

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	mainApi := api.New()

	server := http.Server{
		Addr:    ":8080",
		Handler: mainApi.Handler(),
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		fmt.Printf("Received signal: %v, initiating graceful shutdown\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not gracefully shut down the server: %w", err)
		}

		fmt.Println("Server gracefully stopped")
		return nil
	}
}
