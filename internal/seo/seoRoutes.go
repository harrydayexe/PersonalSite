// Package seo serves the crawler-facing files that live at the origin root:
// robots.txt, a sitemap index, and a sitemap for the hand-written pages.
//
// The blog's own sitemap is produced by GoBlog and served by its handler under
// the blog root, so it is only referenced from here, never rebuilt.
package seo

import (
	"context"
	"encoding/xml"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// staticPaths are the pages rendered outside GoBlog, in the order they should
// appear in the sitemap. They are listed literally rather than discovered from
// the mux because net/http exposes no way to enumerate registered patterns;
// keep this in sync with internal/home and internal/songlinkr.
var staticPaths = []string{
	"/",
	"/songlinkr",
	"/songlinkr/privacy",
	"/songlinkr/support",
}

// PagesSitemapPath is where the sitemap for the non-blog pages is served. It
// is referenced by the sitemap index at SitemapIndexPath.
const PagesSitemapPath = "/pages-sitemap.xml"

// SitemapIndexPath is the sitemap index advertised to crawlers by robots.txt.
// It points at the pages sitemap and at GoBlog's blog sitemap.
const SitemapIndexPath = "/sitemap.xml"

// RobotsPath is where robots.txt is served. Crawlers only ever look for it at
// the origin root.
const RobotsPath = "/robots.txt"

// urlSet is the sitemaps.org 0.9 <urlset> document. No <changefreq> or
// <priority>: Google ignores both.
type urlSet struct {
	XMLName xml.Name `xml:"urlset"`
	NS      string   `xml:"xmlns,attr"`
	URLs    []urlRef `xml:"url"`
}

type urlRef struct {
	Loc string `xml:"loc"`
}

// sitemapIndex is the sitemaps.org 0.9 <sitemapindex> document.
type sitemapIndex struct {
	XMLName  xml.Name     `xml:"sitemapindex"`
	NS       string       `xml:"xmlns,attr"`
	Sitemaps []sitemapRef `xml:"sitemap"`
}

type sitemapRef struct {
	Loc string `xml:"loc"`
}

const sitemapNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

// AddSEORoutes registers GET /robots.txt, GET /sitemap.xml and
// GET /pages-sitemap.xml on the provided mux.
//
// blogSitemapPath is the path GoBlog's handler serves its sitemap at (see
// blog.SitemapPath); it is listed in the sitemap index alongside the pages
// sitemap. Both documents and robots.txt are built once here and served from
// memory. It is not safe for concurrent use during setup, but the resulting
// handlers are.
//
// Neither sitemap emits <lastmod> for the hand-written pages: their content
// changes only when the templates do, and there is no honest timestamp to
// report. GoBlog emits <lastmod> for posts in its own sitemap.
func AddSEORoutes(ctx context.Context, mux *http.ServeMux, logger *slog.Logger, siteURL string, blogSitemapPath string) error {
	logger.DebugContext(ctx, "adding seo routes")

	base := strings.TrimSuffix(siteURL, "/")
	if base == "" {
		return fmt.Errorf("siteURL must not be empty")
	}

	pages := urlSet{NS: sitemapNS}
	for _, p := range staticPaths {
		pages.URLs = append(pages.URLs, urlRef{Loc: base + p})
	}
	pagesXML, err := marshalSitemap(pages)
	if err != nil {
		return err
	}

	index := sitemapIndex{NS: sitemapNS, Sitemaps: []sitemapRef{
		{Loc: base + PagesSitemapPath},
		{Loc: base + blogSitemapPath},
	}}
	indexXML, err := marshalSitemap(index)
	if err != nil {
		return err
	}

	robots := []byte(strings.Join([]string{
		"User-agent: *",
		"Allow: /",
		"",
		"Sitemap: " + base + SitemapIndexPath,
		"",
	}, "\n"))

	serve := func(contentType string, body []byte) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			if _, err := w.Write(body); err != nil {
				logger.ErrorContext(r.Context(), "failed to write seo response", slog.String("path", r.URL.Path), slog.String("error", err.Error()))
			}
		}
	}

	mux.HandleFunc("GET "+RobotsPath, serve("text/plain; charset=utf-8", robots))
	mux.HandleFunc("GET "+SitemapIndexPath, serve("application/xml; charset=utf-8", indexXML))
	mux.HandleFunc("GET "+PagesSitemapPath, serve("application/xml; charset=utf-8", pagesXML))

	logger.DebugContext(ctx, "finished adding seo routes")
	return nil
}

// marshalSitemap renders doc as an XML document with the standard declaration
// prepended, matching the form GoBlog writes its own sitemap in.
func marshalSitemap(doc any) ([]byte, error) {
	body, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(append([]byte(xml.Header), body...), '\n'), nil
}
