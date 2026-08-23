package log

import (
	"log/slog"
	"strings"
)

const LevelPanic = slog.LevelError + 2
const LevelFatal = slog.LevelError + 4

// ParseLevel parses a level string into a slog.Level value.
// Get custom panic/fatal or default levels.
// If the string is not a valid level, it returns slog.LevelInfo.
func ParseLevel(s string) slog.Level {
	var lvl slog.Level

	switch strings.ToUpper(s) {
	case "PANIC":
		lvl = LevelPanic
	case "FATAL":
		lvl = LevelFatal
	}

	_ = lvl.UnmarshalText([]byte(s))

	return lvl
}
