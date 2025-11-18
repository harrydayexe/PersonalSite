# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based personal website that combines static content serving with a dynamic blog system. The blog functionality is powered by the [GoBlog](https://github.com/harrydayexe/GoBlog) framework, which handles markdown post rendering, tag organization, search, and HTMX-powered dynamic updates.

## Architecture

The application has a simple three-layer architecture:

1. **Main Router** (`main.go`): Uses Go's standard library `net/http.ServeMux` for routing
   - Routes `/static/*` to static file server
   - Routes `/blog/*` to the GoBlog server instance (external dependency)
   - Routes `/` to serve the homepage from static content

2. **Static Content**: HTML/CSS/JS files served directly from the `static/` directory

3. **Blog System**: The GoBlog server is initialized with `server.Config` pointing to the `posts/` directory, where markdown files with YAML frontmatter are stored

## Development Commands

### Running Locally
```bash
go run main.go
```

### Docker Development
```bash
# Build and run with Docker Compose (recommended)
docker-compose up --build

# Manual Docker build
docker build -t personal-site .
docker run -p 8080:8080 personal-site
```

### Dependency Management
```bash
# Download dependencies
go mod download

# Update dependencies
go get -u ./...
go mod tidy
```

## Configuration

The application uses environment variables with fallback defaults (configured in `main.go:59-64`):
- `PORT` - HTTP server port (default: 8080)
- `STATIC_DIR` - Static content directory (default: ./static)
- `POSTS_DIR` - Blog posts directory (default: ./posts)

These can be set via docker-compose.yml or exported in the shell for local development.

## Blog Post Format

Blog posts live in `posts/` and use Markdown with YAML frontmatter:
```markdown
---
title: "Your Post Title"
date: 2025-11-18
author: "Your Name"
tags: ["tag1", "tag2"]
slug: "your-post-slug"
---

# Content here...
```

The GoBlog framework handles parsing and rendering these posts.

## Routing Structure

- `/` - Homepage (serves `static/index.html` explicitly, returns 404 for other paths)
- `/blog/*` - All blog routes handled by GoBlog server (stripped prefix)
- `/static/*` - Static file serving (stripped prefix)

Note: The root handler explicitly checks `r.URL.Path != "/"` to prevent serving as a catch-all.

## Docker Build

The Dockerfile uses a multi-stage build:
1. **Builder stage**: Compiles Go binary with CGO disabled for Alpine compatibility
2. **Runtime stage**: Minimal Alpine image with non-root user, copies binary and content directories

Volumes are mounted in docker-compose.yml for live editing of static content and posts without rebuilds.

## GitHub Actions

This repository uses Claude Code Review via GitHub Actions. The workflow (`.github/workflows/claude-code-review.yml`) runs on PRs and uses this CLAUDE.md file to provide context for code reviews.
- Ensure that after all changes, go fmt, go vet and golangci-lint is ran and all warnings are dealt with where possible