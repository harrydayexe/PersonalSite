package main

import (
	"context"
	"embed"
	"log"
	"log/slog"
	"net/http"

	"github.com/harrydayexe/GoWebUtilities/server"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	ctx := context.Background()

	// setDefaultLogger(serverCfg)
	logger := slog.Default()
	logger.Debug("Default logger configured")

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)
	logger.Debug("Static routes added to mux")

	if err := server.Run(ctx, mux); err != nil {
		log.Fatal(err)
	}
}

// // setDefaultLogger sets the default slog logger to be used in the application
// func setDefaultLogger(cfg config.ServerConfig) {
// 	var logger *slog.Logger
// 	var handlerOptions slog.HandlerOptions
//
// 	if cfg.VerboseMode {
// 		handlerOptions = slog.HandlerOptions{Level: slog.LevelDebug}
// 	} else {
// 		handlerOptions = slog.HandlerOptions{Level: slog.LevelInfo}
// 	}
//
// 	if cfg.Environment == config.Local {
// 		logger = slog.New(slog.NewTextHandler(os.Stdout, &handlerOptions))
// 	} else {
// 		logger = slog.New(slog.NewJSONHandler(os.Stdout, &handlerOptions))
// 	}
//
// 	slog.SetDefault(logger)
// }
