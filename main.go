package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/harrydayexe/GoWebUtilities/middleware"
	"github.com/harrydayexe/GoWebUtilities/server"
	blogcontent "github.com/harrydayexe/PersonalSite/internal/blog"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed posts
var postsFiles embed.FS

//go:embed templates
var templateFiles embed.FS

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)

	postsFS, err := fs.Sub(postsFiles, "posts")
	if err != nil {
		log.Fatal(err)
	}

	templatesFS, err := fs.Sub(templateFiles, "templates")
	if err != nil {
		log.Fatal(err)
	}

	if err := blogcontent.AddBlogRoutes(ctx, mux, postsFS, templatesFS, logger); err != nil {
		log.Fatal(err)
	}

	stack := middleware.CreateStack(
		middleware.NewLoggingMiddleware(logger),
	)

	if err := server.Run(ctx, stack(mux)); err != nil {
		log.Fatal(err)
	}
}
