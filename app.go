package droid

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/novoseltsev-stan/droid/log"
	"golang.org/x/sync/errgroup"
)

const defaultShutdownTimeout = 5 * time.Second

type Instance interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type App struct {
	instances       []Instance
	shutdownTimeout time.Duration
	logger          *slog.Logger
}

func New(opts ...Option) *App {
	app := &App{
		shutdownTimeout: defaultShutdownTimeout,
		instances:       make([]Instance, 0),
	}

	for _, opt := range opts {
		opt(app)
	}

	if app.logger != nil {
		slog.SetDefault(app.logger)
	}

	return app
}

func (a *App) Run() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	select {
	case sig := <-ch:
		slog.Info("Interrupted by signal", slog.String("signal", sig.String()))
	case err := <-a.start():
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start server", log.Err(err))
		}
	}

	if err := a.stop(); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("Failed to gracefull shutdown", log.Err(err))
		return err
	}

	return nil
}

func (a *App) start() <-chan error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, inst := range a.instances {
		eg.Go(func() error {
			return inst.Start(ctx)
		})
	}

	resCh := make(chan error)
	go func() {
		defer close(resCh)
		resCh <- eg.Wait()
	}()

	return resCh
}

func (a *App) stop() error {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	eg, ctx := errgroup.WithContext(timeoutCtx)
	for _, inst := range a.instances {
		eg.Go(func() error {
			return inst.Stop(ctx)
		})
	}

	return eg.Wait()
}
