package main

import (
	"go-backend/routers"
	"log"
	"net/http"
)

func main() {
	const port string = ":8080"
	r := routers.Routers()

	log.Printf("Serving on Port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))

}
