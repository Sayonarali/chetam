package factory

import (
	"log/slog"
	"os"
)

func provideSlog() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
