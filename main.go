package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"

	"github.com/harrydayexe/GoWebUtilities/server"
	blogcontent "github.com/harrydayexe/PersonalSite/internal/blog"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed posts
var postsFiles embed.FS

func main() {
	ctx := context.Background()
	logger := slog.Default()

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)

	postsFS, err := fs.Sub(postsFiles, "posts")
	if err != nil {
		log.Fatal(err)
	}

	if err := blogcontent.AddBlogRoutes(ctx, mux, postsFS, logger); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(ctx, mux); err != nil {
		log.Fatal(err)
	}
}
