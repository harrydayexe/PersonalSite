package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/harrydayexe/GoWebUtilities/server"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()
	staticcontent.AddStaticRoutes(mux, staticFiles)

	if err := server.Run(ctx, mux); err != nil {
		log.Fatal(err)
	}
}
