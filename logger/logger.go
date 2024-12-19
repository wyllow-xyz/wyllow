package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var logger *slog.Logger

const LevelFatal slog.Level = slog.LevelError + 4

// Logging functions for predefined log levels.
//
// Debug: Detailed messages for debugging purposes.
// Info: Informational messages.
// Warn: Warnings about potential issues or unusual conditions.
// Error: Error messages indicating an issue that needs attention.
// Fatal: Critical errors that terminate the application.
var (
	Debug = makeLogFunc(slog.LevelDebug)
	Info  = makeLogFunc(slog.LevelInfo)
	Warn  = makeLogFunc(slog.LevelWarn)
	Error = makeLogFunc(slog.LevelError)
	Fatal = makeLogFunc(LevelFatal)
)

// Initializes the global logger and set up logging to stderr with a default level of debug.
func init() {
	w := os.Stderr
	logger = slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					switch a.Value.Any() {
					case LevelFatal:
						return slog.Attr{
							Key:   slog.LevelKey,
							Value: slog.StringValue("\x1b[91mFATAL\x1b[0m"),
						}
					}
				}
				return a
			},
		}),
	)
	slog.SetDefault(logger)
}

// Returns a logging function for a specific log level.
// The returned function formats the log message using fmt.Sprintf if arguments are provided.
// If the level is LevelFatal, the function terminates the application with a call to os.Exit(1).
func makeLogFunc(level slog.Level) func(msg string, args ...any) {
	return func(msg string, args ...any) {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args...)
		}
		logger.Log(nil, level, msg)

		if level == LevelFatal {
			os.Exit(1)
		}
	}
}
