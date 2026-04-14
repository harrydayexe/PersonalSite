// Package blog provides HTTP route registration for the blog feed.
package blog

import (
	"context"
	"embed"
	"io/fs"
	"log/slog"
	"net/http"

	goblogconfig "github.com/harrydayexe/GoBlog/v2/pkg/config"
	"github.com/harrydayexe/GoBlog/v2/pkg/generator"
	goblogserver "github.com/harrydayexe/GoBlog/v2/pkg/server"
)

//go:embed all:templates
var templateFiles embed.FS

// AddBlogRoutes registers the blog HTTP routes at /blog/ on the provided mux.
// It uses custom templates matching the site's visual identity and serves posts
// from the given fs.FS. It is not safe for concurrent use during setup, but the
// resulting handler is.
func AddBlogRoutes(ctx context.Context, mux *http.ServeMux, posts fs.FS, logger *slog.Logger) error {
	templatesFS, err := fs.Sub(templateFiles, "templates")
	if err != nil {
		return err
	}

	renderer, err := generator.NewTemplateRenderer(templatesFS)
	if err != nil {
		return err
	}

	gen := generator.New(
		posts,
		renderer,
		goblogconfig.WithBaseOption(goblogconfig.WithBlogRoot("/blog/")),
		goblogconfig.WithSiteTitle("Harry Day{}"),
	)

	blog, err := gen.Generate(ctx)
	if err != nil {
		return err
	}

	handler := goblogserver.Handler(blog, logger, goblogconfig.WithBlogRoot("/blog/"))
	mux.Handle("/blog/", http.StripPrefix("/blog", handler))
	return nil
}
