package main

// Need to implement a web server. new http.ServeMux

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port string = ":8000"
	fmt.Println("Hello World")
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.Handle("/hello_world", http.HandlerFunc(helloWorldHandler))

	srv := &http.Server{
		Addr:    port,
		Handler: corsMux,
	}
	log.Printf("Serving on Port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
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
