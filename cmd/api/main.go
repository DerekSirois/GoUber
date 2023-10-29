package main

import (
	"GoUber/pkg/db"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err := db.InitDb()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Starting on port 8080")
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
