// Package home provides the HTTP handler for the site homepage.
package home

import (
	"context"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/harrydayexe/GoBlog/v2/pkg/models"
	goblogparser "github.com/harrydayexe/GoBlog/v2/pkg/parser"
)

const blogRoot = "/blog/"

type homePageData struct {
	Posts       models.PostList
	Year        int
	BlogRoot    string
	Environment string
	SiteURL     string
}

// AddHomeRoute registers the GET /{$} handler for the site homepage on the provided mux.
// It parses posts once at startup using GoBlog's parser and serves the home.tmpl template.
// It is not safe for concurrent use during setup, but the resulting handler is.
func AddHomeRoute(ctx context.Context, mux *http.ServeMux, postsFS fs.FS, templatesFS fs.FS, logger *slog.Logger, environment string, siteURL string) error {
	tmpl, err := template.ParseFS(templatesFS, "pages/home.tmpl")
	if err != nil {
		return err
	}

	posts, err := goblogparser.New().ParseDirectory(ctx, postsFS)
	if err != nil {
		logger.WarnContext(ctx, "some posts failed to parse for home route", slog.String("error", err.Error()))
	}
	posts.SortByDate()
	if len(posts) > 10 {
		posts = posts[:10]
	}

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, homePageData{
			Posts:       posts,
			Year:        time.Now().Year(),
			BlogRoot:    blogRoot,
			Environment: environment,
			SiteURL:     siteURL,
		}); err != nil {
			logger.ErrorContext(r.Context(), "failed to render homepage", slog.String("error", err.Error()))
		}
	})

	return nil
}
