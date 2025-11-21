package main

import (
	"log"
	"net/http"
	"os"

	"github.com/harrydayexe/GoBlog/pkg/server"
)

func main() {
	// Get configuration from environment variables with defaults
	port := getEnv("PORT", "8080")
	staticDir := getEnv("STATIC_DIR", "./static")
	postsDir := getEnv("POSTS_DIR", "./posts")

	// Create a new ServeMux for routing
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Initialize GoBlog server for blog functionality
	blogOpts := server.Options{
		ContentPath:  postsDir,
		EnableCache:  true, // Use Ristretto caching for better performance
		EnableSearch: true, // Enable Bleve full-text search
		PostsPerPage: 10,   // Number of posts per page
	}

	blogServer, err := server.New(blogOpts)
	if err != nil {
		log.Fatalf("Failed to initialize blog server: %v", err)
	}
	defer func() {
		if err := blogServer.Close(); err != nil {
			log.Printf("Error closing blog server: %v", err)
		}
	}()

	// Attach blog routes to the mux at /blog path
	blogServer.AttachRoutes(mux, "/blog")

	// Home page handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, staticDir+"/index.html")
	})

	// Start the server
	addr := ":" + port
	log.Printf("Server starting on %s", addr)
	log.Printf("Static content directory: %s", staticDir)
	log.Printf("Blog posts directory: %s", postsDir)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// getEnv retrieves an environment variable or returns a default value if not set.
// This is used to configure the server from the environment while providing
// sensible defaults for local development.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
