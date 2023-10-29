package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	log.Println("Starting on port 8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
