# Personal Site

A dynamic personal website built with Go, featuring static content serving and a dynamic blog powered by the [GoBlog](https://github.com/harrydayexe/GoBlog) framework.

## Project Structure

```
PersonalSite/
├── main.go                 # Main application entry point
├── go.mod                  # Go module definition
├── Dockerfile              # Docker build configuration
├── docker-compose.yml      # Docker Compose configuration
├── static/                 # Static content directory
│   ├── index.html         # Homepage
│   └── about.html         # About page
├── posts/                  # Blog posts directory (Markdown files)
│   └── examples/
│       └── welcome.md     # Example blog post
└── templates/              # HTML templates (for future use)
```

## Features

- **Static Content Serving**: Serves static HTML, CSS, and JavaScript files from the `static/` directory
- **Dynamic Blog**: Markdown-based blog system powered by GoBlog
  - Automatic rendering of markdown posts
  - Tag-based organization
  - Real-time search capabilities
  - HTMX-powered dynamic updates
- **Standard Library Router**: Uses Go's built-in `net/http` router
- **Docker Support**: Containerized deployment with multi-stage builds
- **Environment Configuration**: Configurable via environment variables

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerized deployment)

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **Access the site:**
   - Homepage: http://localhost:8080
   - Blog: http://localhost:8080/blog/
   - Static content: http://localhost:8080/static/

### Docker Deployment

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

2. **Or build and run manually:**
   ```bash
   docker build -t personal-site .
   docker run -p 8080:8080 personal-site
   ```

## Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Port number for the HTTP server |
| `STATIC_DIR` | `./static` | Directory for static content |
| `POSTS_DIR` | `./posts` | Directory for blog posts |

## Creating Blog Posts

Blog posts are written in Markdown with YAML frontmatter. Create new posts in the `posts/` directory:

```markdown
---
title: "Your Post Title"
date: 2025-11-18
author: "Your Name"
tags: ["tag1", "tag2"]
slug: "your-post-slug"
---

# Your Post Title

Your content here...
```

## Routes

- `/` - Homepage
- `/blog/` - Blog listing and posts
- `/static/` - Static content (CSS, JS, images, etc.)

## Technology Stack

- **Language**: Go
- **Framework**: GoBlog (for blog functionality)
- **Router**: Go standard library (`net/http`)
- **Container**: Docker
- **Blog Format**: Markdown with YAML frontmatter

## License

MIT
