package main

import (
	"log/slog"
	"os"

	"myoffice-test-task/internal/app"
	"myoffice-test-task/internal/config"
)

func main() {
	// parse config
	cfg, err := config.Parse()
	if err != nil {
		slog.Error("failed to parse config",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	// run app
	err = app.Run(cfg)
	if err != nil {
		slog.Error("failed to run app.Run",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}
