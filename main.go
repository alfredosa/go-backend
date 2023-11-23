package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	const port string = ":8000"
	const staticDir string = "./static/"

	api_config := apiConfig{fileServerHitCount: 0}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewareCors)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, staticDir))
	FileServer(r, "/app", filesDir, api_config)

	r.Get("/healthz", http.HandlerFunc(healthHandler))
	r.Get("/metrics", http.HandlerFunc(api_config.metricsHandler))
	r.Get("/reset", http.HandlerFunc(api_config.resetHandler))

	log.Printf("Serving on Port: %s\n", port)
	http.ListenAndServe(port, r)

}

func FileServer(r chi.Router, path string, root http.FileSystem, api_config apiConfig) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"
	log.Printf("ADDING CONFIG GET")
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := api_config.incrementFileServerHitCount(http.StripPrefix(pathPrefix, http.FileServer(root)))
		fs.ServeHTTP(w, r)
	})
}
