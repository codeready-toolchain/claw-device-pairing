package logger

import (
	"log/slog"
	"os"
)

// Init initializes the structured logger with JSON output
func Init() {
	// Determine log level based on ENV variable
	logLevel := slog.LevelInfo
	if os.Getenv("ENV") == "development" {
		logLevel = slog.LevelDebug
	}

	// Create JSON handler
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})

	// Set as default logger
	slog.SetDefault(slog.New(handler))
}
