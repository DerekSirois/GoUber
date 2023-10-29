package main

import (
	"GoUber/pkg/handler"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")

	return router
}
