package log

import (
	"fmt"
	"log/slog"
	"strings"
)

const (
	DebugLevel = slog.LevelDebug
	InfoLevel  = slog.LevelInfo
	WarnLevel  = slog.LevelWarn
	ErrorLevel = slog.LevelError
)

func ParseLevel(lvl string) (slog.Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	default:
		return 0, fmt.Errorf("unrecognized level: %s", lvl)
	}
}
