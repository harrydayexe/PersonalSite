// Package main implements a personal website combining static content serving
// with a dynamic blog system powered by the GoBlog framework.
//
// The application serves three types of content:
//   - Static HTML/CSS/JS files from the /static directory
//   - A dynamic blog system under /blog using the GoBlog framework
//   - A homepage served from static/index.html
//
// # Architecture
//
// The application uses a simple three-layer architecture:
//
//  1. Main Router (main.go): Uses Go's standard library http.ServeMux for routing
//     - Routes /static/* to static file server
//     - Routes /blog/* to the GoBlog server instance
//     - Routes / to serve the homepage from static content
//
//  2. Static Content: HTML/CSS/JS files served directly from the static/ directory
//
//  3. Blog System: The GoBlog server handles markdown post rendering, tag organization,
//     search functionality, and HTMX-powered dynamic updates
//
// # Configuration
//
// The server is configured via environment variables with fallback defaults:
//   - PORT: HTTP server port (default: 8080)
//   - STATIC_DIR: Static content directory (default: ./static)
//   - POSTS_DIR: Blog posts directory (default: ./posts)
//
// These can be set via docker-compose.yml or exported in the shell for local development.
//
// # Blog Posts
//
// Blog posts are markdown files in the posts/ directory with YAML frontmatter:
//
//	---
//	title: "Your Post Title"
//	date: 2025-11-18
//	author: "Your Name"
//	tags: ["tag1", "tag2"]
//	slug: "your-post-slug"
//	---
//
//	# Content here...
//
// The GoBlog framework handles parsing and rendering these posts with support for:
//   - Markdown rendering with syntax highlighting
//   - Full-text search via Bleve
//   - Tag-based organization and filtering
//   - HTMX-powered dynamic updates
//   - RSS/Atom feeds and sitemaps
//
// # Docker Deployment
//
// The Dockerfile uses a multi-stage build:
//  1. Builder stage: Compiles Go binary with CGO disabled for Alpine compatibility
//  2. Runtime stage: Minimal Alpine image with non-root user
//
// Run with Docker Compose:
//
//	docker-compose up --build
//
// Or manually:
//
//	docker build -t personal-site .
//	docker run -p 8080:8080 personal-site
//
// See [github.com/harrydayexe/GoBlog] for blog framework documentation.
package main
