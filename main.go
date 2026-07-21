package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/harrydayexe/GoWebUtilities/config"
	"github.com/harrydayexe/GoWebUtilities/logging"
	"github.com/harrydayexe/GoWebUtilities/middleware"
	"github.com/harrydayexe/GoWebUtilities/server"
	blogcontent "github.com/harrydayexe/PersonalSite/internal/blog"
	homecontent "github.com/harrydayexe/PersonalSite/internal/home"
	"github.com/harrydayexe/PersonalSite/internal/songlinkr"
	staticcontent "github.com/harrydayexe/PersonalSite/internal/static-content"
)

type appConfig struct {
	config.ServerConfig
	SiteURL string `env:"SITE_URL" envDefault:"https://harryday.dev"`
}

func (c appConfig) Validate() error {
	if err := c.ServerConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.SiteURL) == "" {
		return fmt.Errorf("SITE_URL must not be empty")
	}
	return nil
}

//go:embed static/*
var staticFiles embed.FS

//go:embed posts
var postsFiles embed.FS

//go:embed templates
var templateFiles embed.FS

func main() {
	ctx := context.Background()
	cfg, err := config.ParseConfig[appConfig]()
	if err != nil {
		log.Fatalf("failed to create config from environment: %s", err.Error())
	}

	logging.SetDefaultLogger(cfg.ServerConfig)
	logger := slog.Default()

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

	if cfg.Environment == config.Local {
		mux.HandleFunc("GET /reload", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			w.WriteHeader(http.StatusOK)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			<-r.Context().Done()
		})
	}

	if err := homecontent.AddHomeRoute(ctx, mux, postsFS, templatesFS, logger, string(cfg.Environment), cfg.SiteURL); err != nil {
		log.Fatal(err)
	}

	if err := blogcontent.AddBlogRoutes(ctx, mux, postsFS, templatesFS, logger, string(cfg.Environment), cfg.SiteURL); err != nil {
		log.Fatal(err)
	}

	if err := songlinkr.AddSonglinkrRoutes(ctx, mux, templatesFS, logger, string(cfg.Environment), cfg.SiteURL); err != nil {
		log.Fatal(err)
	}

	stack := middleware.CreateStack(
		middleware.NewLoggingMiddleware(logger),
	)

	if err := server.Run(ctx, stack(mux)); err != nil {
		log.Fatal(err)
	}
}
