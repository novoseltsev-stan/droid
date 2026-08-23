package log

import (
	"fmt"
	"log/slog"
	"runtime/debug"
)

// Err returns an attr with the given error message under the "error" key.
func Err(err error) slog.Attr {
	if err == nil {
		return slog.String("error", "")
	}

	return slog.String("error", err.Error())
}

// Panic returns an attr with the given panic value under the "panic" key.
func Panic(p any) slog.Attr {
	return slog.String("panic", fmt.Sprint(p))
}

// StackTrace returns an attr with the current stack trace under the "stack_trace" key.
func StackTrace() slog.Attr {
	return slog.String("stack_trace", string(debug.Stack()))
}

// AppVersion returns an attr with the application version under the "app_version" key.
func AppVersion() slog.Attr {
	bi, ok := debug.ReadBuildInfo()
	if !ok || bi == nil {
		return slog.String("app_version", "")
	}

	return slog.String("app_version", bi.Main.Version)
}
