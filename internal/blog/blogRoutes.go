// Package blog provides HTTP route registration for the blog feed.
package blog

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"

	goblogconfig "github.com/harrydayexe/GoBlog/v2/pkg/config"
	"github.com/harrydayexe/GoBlog/v2/pkg/generator"
	goblogserver "github.com/harrydayexe/GoBlog/v2/pkg/server"
)

// AddBlogRoutes registers the blog HTTP routes at /blog/ on the provided mux.
// It uses custom templates matching the site's visual identity and serves posts
// from the given fs.FS. It is not safe for concurrent use during setup, but the
// resulting handler is.
func AddBlogRoutes(ctx context.Context, mux *http.ServeMux, posts fs.FS, templates fs.FS, logger *slog.Logger) error {
	logger.DebugContext(ctx, "adding blog routes")

	renderer, err := generator.NewTemplateRenderer(templates)
	if err != nil {
		return err
	}

	gen := generator.New(
		posts,
		renderer,
		goblogconfig.WithBaseOption(goblogconfig.WithBlogRoot("/blog/")),
		goblogconfig.WithSiteTitle("Harry Day{}"),
	)

	logger.DebugContext(ctx, "generator created", slog.String("config", gen.String()))

	blog, err := gen.Generate(ctx)
	if err != nil {
		return err
	}

	if len(blog.Index) == 0 {
		return fmt.Errorf("blog generation produced empty index: check template structure")
	}

	logger.DebugContext(ctx, "blog generated", slog.String("index", string(blog.Index)))

	handler := goblogserver.Handler(blog, logger, goblogconfig.WithBlogRoot("/blog/"))
	mux.Handle("/blog/", handler)

	logger.DebugContext(ctx, "finished adding blog routes")
	return nil
}
