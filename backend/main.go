package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define the directory for the static files
	staticDir := "../frontend/build"

	// Create a file server for the static files
	fs := http.FileServer(http.Dir(staticDir))

	// Wrap the file server to serve index.html when the file is not found
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)

		// Check if the request resulted in a 404 (Not Found) status
		if w.Header().Get("X-Content-Type-Options") == "nosniff" {
			// If so, serve the index.html file
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
		}
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
