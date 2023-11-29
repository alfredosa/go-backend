package routers

import (
	"net/http"
	"os"
	"path/filepath"

	"go-backend/handlers"
	"go-backend/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Routers(db *sqlx.DB) *chi.Mux {
	const staticDir string = "./static/"

	api_config := handlers.ApiConfig{
		FileServerHitCount: 0,
		DB:                 db,
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.MiddlewareCors)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, staticDir))
	api_config.FileServer(r, "/app", filesDir, &api_config)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", api_config.HealthHandler)
	apiRouter.Get("/reset", api_config.ResetHandler)
	apiRouter.Post("/validate_chirp", api_config.ValidateChirpHandler)
	r.Mount("/api", apiRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", api_config.MetricsHandler)
	r.Mount("/admin", adminRouter)
	return r
}
