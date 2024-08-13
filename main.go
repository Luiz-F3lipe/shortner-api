package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Luiz-F3lipe/shortner-api/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to excute server", "error", err)
		return
	}

	slog.Info("All systems offline")
}

func run() error {
	handler := api.NewHandler()

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
