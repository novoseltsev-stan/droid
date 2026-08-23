package log

import (
	"log/slog"
	"strings"
)

const LevelFatal = slog.LevelError + 4

// ParseLevel parses a level string into a slog.Level value.
// Get fatal level for "FATAL" string, otherwise use slog.Level.UnmarshalText.
// If the string is not a valid level, it returns slog.LevelInfo.
func ParseLevel(s string) slog.Level {
	if strings.EqualFold(strings.ToUpper(s), "FATAL") {
		return LevelFatal
	}

	var lvl slog.Level

	_ = lvl.UnmarshalText([]byte(s))

	return lvl
}
