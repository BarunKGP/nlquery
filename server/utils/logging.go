package utils

import (
	"io"
	"log/slog"
	"os"
)

func CreateLogger() *slog.Logger {

	var writer io.Writer

	// Add a MultiWriter if LOG_FILE is defined
	le, ok := os.LookupEnv("LOG_FILE")
	if ok {
		file, err := os.OpenFile(le, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		writer = io.MultiWriter(file, os.Stderr)
	} else {
		// Write to stderr by default
		writer = os.Stderr
	}

	handlerOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(writer, handlerOpts))

	slog.SetDefault(logger)

	return logger
}
