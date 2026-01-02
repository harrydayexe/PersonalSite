package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)

	if err := Run(ctx, mux, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run starts the HTTP server with the provided handler.
func Run(ctx context.Context, srv http.Handler, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
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
