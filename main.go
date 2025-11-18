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
	blogConfig := server.Config{
		PostsDirectory: postsDir,
		// Add additional configuration as needed
	}

	blogServer, err := server.New(blogConfig)
	if err != nil {
		log.Fatalf("Failed to initialize blog server: %v", err)
	}

	// Mount blog routes under /blog
	mux.Handle("/blog/", http.StripPrefix("/blog", blogServer))

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

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
