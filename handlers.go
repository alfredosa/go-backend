package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileServerHitCount int
}

func (a *apiConfig) incrementFileServerHitCount(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.fileServerHitCount++
		log.Printf("printing COUNT: %d", a.fileServerHitCount)
		next.ServeHTTP(w, r)
	})
}

func (a *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Printf("HITS: %d", a.fileServerHitCount)

	html := `
    <html>
    <body>
        <h1>Welcome, Chirpy Admin</h1>
        <p>Chirpy has been visited %d times!</p>
    </body>
    </html>
    `
	w.Write([]byte(fmt.Sprintf(html, a.fileServerHitCount)))

}

func (a *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	a.fileServerHitCount = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func FileServer(r chi.Router, path string, root http.FileSystem, api_config *apiConfig) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := api_config.incrementFileServerHitCount(http.StripPrefix(pathPrefix, http.FileServer(root)))
		fs.ServeHTTP(w, r)
	})
}

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type validatePost struct {
		Body string
	}

	var response validatePost
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(response.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"valid":true`))

}
