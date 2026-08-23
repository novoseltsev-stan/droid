package droid

import (
	"log/slog"
	"time"
)

type Option func(*App)

func Instances(inst ...Instance) Option {
	return func(a *App) { a.instances = append(a.instances, inst...) }
}

func ShutdownTimeout(d time.Duration) Option {
	return func(a *App) { a.shutdownTimeout = d }
}

func Logger(l *slog.Logger) Option {
	return func(a *App) { a.logger = l }
}
