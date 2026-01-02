package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/harrydayexe/PersonalSite/internal/config"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	ctx := context.Background()
	var serverCfg config.ServerConfig = parseConfig[config.ServerConfig]()

	setDefaultLogger(serverCfg)

	logger := slog.Default()
	logger.Debug("Default logger configured")

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)
	logger.Debug("Static routes added to mux")

	if err := run(ctx, mux, serverCfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// parseConfig sets a config object based on environment variables
func parseConfig[C any]() C {
	var cfg C

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

// setDefaultLogger sets the default slog logger to be used in the application
func setDefaultLogger(cfg config.ServerConfig) {
	var logger *slog.Logger
	var handlerOptions slog.HandlerOptions

	if cfg.VerboseMode {
		handlerOptions = slog.HandlerOptions{Level: slog.LevelDebug}
	} else {
		handlerOptions = slog.HandlerOptions{Level: slog.LevelInfo}
	}

	if cfg.Environment == config.Local {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &handlerOptions))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &handlerOptions))
	}

	slog.SetDefault(logger)
}

// Run starts the HTTP server with the provided handler.
func run(
	ctx context.Context,
	srv http.Handler,
	cfg config.ServerConfig,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: srv,
	}
	go func() {
		logger.Info(
			"server listening",
			slog.String("address", httpServer.Addr),
		)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
