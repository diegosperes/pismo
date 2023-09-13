package util

import (
	"log/slog"
	"os"
)

func ConfigureDefaultLogger(levelName string) error {
	level := new(slog.Level)

	if levelErr := level.UnmarshalText([]byte(levelName)); levelErr != nil {
		return levelErr
	}

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{Level: level},
		),
	)

	slog.SetDefault(logger)
	return nil
}
