package main

import (
	"fmt"
	"log"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", a.fileServerHitCount)))
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
