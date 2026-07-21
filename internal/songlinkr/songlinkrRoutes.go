// Package songlinkr provides HTTP handlers for the SongLinkr portfolio pages.
package songlinkr

import (
	"context"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"time"
)

const blogRoot = "/blog/"

type songlinkrPageData struct {
	Year        int
	BlogRoot    string
	Environment string
	SiteURL     string
}

// AddSonglinkrRoutes registers the SongLinkr landing, privacy, and support
// pages on the provided mux. It parses the songlinkr page templates together
// with the shared header/footer partials once at startup, and serves them at
// GET /songlinkr, GET /songlinkr/privacy, and GET /songlinkr/support.
// Trailing-slash variants are redirected to their canonical form. It is not
// safe for concurrent use during setup, but the resulting handlers are.
func AddSonglinkrRoutes(ctx context.Context, mux *http.ServeMux, templatesFS fs.FS, logger *slog.Logger, environment string, siteURL string) error {
	tmpl, err := template.ParseFS(
		templatesFS,
		"pages/songlinkr.tmpl",
		"pages/songlinkr-privacy.tmpl",
		"pages/songlinkr-support.tmpl",
		"partials/header.tmpl",
		"partials/footer.tmpl",
	)
	if err != nil {
		return err
	}

	data := songlinkrPageData{
		Year:        time.Now().Year(),
		BlogRoot:    blogRoot,
		Environment: environment,
		SiteURL:     siteURL,
	}

	render := func(name string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
				logger.ErrorContext(r.Context(), "failed to render songlinkr page", slog.String("template", name), slog.String("error", err.Error()))
			}
		}
	}

	redirectTo := func(target string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		}
	}

	mux.HandleFunc("GET /songlinkr", render("songlinkr.tmpl"))
	mux.HandleFunc("GET /songlinkr/{$}", redirectTo("/songlinkr"))

	mux.HandleFunc("GET /songlinkr/privacy", render("songlinkr-privacy.tmpl"))
	mux.HandleFunc("GET /songlinkr/privacy/{$}", redirectTo("/songlinkr/privacy"))

	mux.HandleFunc("GET /songlinkr/support", render("songlinkr-support.tmpl"))
	mux.HandleFunc("GET /songlinkr/support/{$}", redirectTo("/songlinkr/support"))

	return nil
}
