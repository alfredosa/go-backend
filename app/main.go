package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	const port string = ":8080"
	const staticDir string = "./static/"

	api_config := apiConfig{
		fileServerHitCount: 0,
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewareCors)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, staticDir))
	FileServer(r, "/app", filesDir, &api_config)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", healthHandler)
	apiRouter.Get("/reset", api_config.resetHandler)
	apiRouter.Post("/validate_chirp", validateChirpHandler)
	r.Mount("/api", apiRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", api_config.metricsHandler)
	r.Mount("/admin", adminRouter)

	log.Printf("Serving on Port: %s\n", port)
	http.ListenAndServe(port, r)

}
