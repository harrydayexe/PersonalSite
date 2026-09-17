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
	// FeedsEnabled mirrors models.BaseData.FeedsEnabled, which the blog's
	// templates get from the generator. The home page is rendered outside the
	// generator, so the field is set here using the same rule GoBlog applies:
	// feeds exist when a base URL is configured.
	FeedsEnabled bool
}

// AddHomeRoute registers the GET /{$} handler for the site homepage on the provided mux.
// It parses posts once at startup using GoBlog's parser and serves the home.tmpl template.
// It is not safe for concurrent use during setup, but the resulting handler is.
func AddHomeRoute(ctx context.Context, mux *http.ServeMux, postsFS fs.FS, templatesFS fs.FS, logger *slog.Logger, environment string, siteURL string) error {
	// schema.tmpl is shared with the blog's head partial so the Person and
	// WebSite JSON-LD nodes stay identical across the two template trees.
	// rss-icon.tmpl is shared for the same reason: one glyph, three trees.
	tmpl, err := template.ParseFS(templatesFS, "pages/home.tmpl", "partials/schema.tmpl", "partials/rss-icon.tmpl")
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
			Posts:        posts,
			Year:         time.Now().Year(),
			BlogRoot:     blogRoot,
			Environment:  environment,
			SiteURL:      siteURL,
			FeedsEnabled: siteURL != "",
		}); err != nil {
			logger.ErrorContext(r.Context(), "failed to render homepage", slog.String("error", err.Error()))
		}
	})

	return nil
}
