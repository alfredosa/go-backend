package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type ApiConfig struct {
	FileServerHitCount int
	DB                 *sqlx.DB
}

func (a *ApiConfig) incrementFileServerHitCount(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.FileServerHitCount++
		log.Printf("printing COUNT: %d", a.FileServerHitCount)
		next.ServeHTTP(w, r)
	})
}

func (a *ApiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Printf("HITS: %d", a.FileServerHitCount)

	html := `
    <html>
    <body>
        <h1>Welcome, Chirpy Admin</h1>
        <p>Chirpy has been visited %d times!</p>
    </body>
    </html>
    `
	w.Write([]byte(fmt.Sprintf(html, a.FileServerHitCount)))

}

func (a *ApiConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {
	a.FileServerHitCount = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset"))
}

func (a *ApiConfig) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (a *ApiConfig) FileServer(r chi.Router, path string, root http.FileSystem, api_config *ApiConfig) {
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

func (a *ApiConfig) ValidateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	RespondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: CleanedBody(params.Body),
	})

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func CleanedBody(payload string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Fields(payload)

	for _, badWord := range badWords {
		for i := range words {
			if strings.EqualFold(words[i], badWord) {
				words[i] = "****"
			}
		}
	}

	return strings.Join(words, " ")
}
