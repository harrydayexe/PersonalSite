package staticcontent

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

func AddStaticRoutes(mux *http.ServeMux, staticFiles embed.FS) {
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", http.FileServer(http.FS(staticFS)))
}
