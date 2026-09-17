// Package blog provides HTTP route registration for the blog feed.
package blog

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"
	"time"

	goblogconfig "github.com/harrydayexe/GoBlog/v2/pkg/config"
	"github.com/harrydayexe/GoBlog/v2/pkg/generator"
	goblogparser "github.com/harrydayexe/GoBlog/v2/pkg/parser"
	goblogserver "github.com/harrydayexe/GoBlog/v2/pkg/server"
)

const blogRoot = "/blog/"

// AddBlogRoutes registers the blog HTTP routes at /blog/ on the provided mux.
// It uses custom templates matching the site's visual identity and serves posts
// from the given fs.FS. It is not safe for concurrent use during setup, but the
// resulting handler is.
func AddBlogRoutes(ctx context.Context, mux *http.ServeMux, posts fs.FS, templates fs.FS, logger *slog.Logger, environment string, siteURL string) error {
	logger.DebugContext(ctx, "adding blog routes")

	renderer, err := generator.NewTemplateRenderer(
		templates,
		goblogconfig.WithFuncs(template.FuncMap{
			"hasPrefix": strings.HasPrefix,
			"isoDate":   func(t time.Time) string { return t.UTC().Format(time.RFC3339) },
			// inc converts a zero-based range index into a one-based
			// schema.org ListItem position.
			"inc": func(i int) int { return i + 1 },
			// wordCount approximates a post's length for schema.org
			// wordCount. It counts whitespace-separated tokens in the raw
			// markdown, so fences, link syntax and punctuation inflate it
			// slightly; the property is a hint to search engines, not a
			// figure displayed anywhere.
			"wordCount": func(s string) int { return len(strings.Fields(s)) },
		}),
	)
	if err != nil {
		return err
	}

	gen := generator.New(
		posts,
		renderer,
		goblogconfig.WithBlogRoot(blogRoot).AsGeneratorOption(),
		goblogconfig.WithSiteTitle("Harry Day"),
		goblogconfig.WithEnvironment(environment),
		goblogconfig.WithCustomData(map[string]any{"siteURL": siteURL}),
	)
	gen.ParserConfig = goblogparser.Config{
		EnableCodeHighlighting: true,
	}

	logger.DebugContext(ctx, "generator created", slog.String("config", gen.String()))

	blog, err := gen.Generate(ctx)
	if err != nil {
		return err
	}

	if len(blog.Index) == 0 {
		return fmt.Errorf("blog generation produced empty index: check template structure")
	}

	logger.DebugContext(ctx, "blog generated", slog.String("index", string(blog.Index)))

	assetsFS, err := fs.Sub(posts, "images")
	if err != nil {
		return err
	}

	handler := goblogserver.Handler(
		blog,
		logger,
		goblogconfig.WithBlogRoot(blogRoot),
		goblogconfig.WithAssetsDir(assetsFS),
	)
	mux.Handle(blogRoot, handler)

	logger.DebugContext(ctx, "finished adding blog routes")
	return nil
}
