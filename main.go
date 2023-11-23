package main

import (
	"log"
	"net/http"
)

func main() {
	const port string = ":8000"
	const staticDir string = "./static/"

	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(staticDir))))
	mux.Handle("/healthz", http.HandlerFunc(healthHandler))
	mux.Handle("/hello_world", http.HandlerFunc(helloWorldHandler))

	srv := &http.Server{
		Addr:    port,
		Handler: corsMux,
	}
	log.Printf("Serving on Port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// example using http.ServeFile and handlerFunc
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// serve files from static folder
	// return 200 status code
	http.ServeFile(w, r, "./static/index.html")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
