package main

import (
	"go-backend/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")

	go func() {
		<-sigs

		db.Close()
		log.Printf("Closed DB Connection")

		log.Printf("Cleanup completed. Exiting...")

		os.Exit(0)
	}()

	const port string = ":8080"
	r := routers.Routers(db)

	log.Printf("Serving on Port: %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}

}
