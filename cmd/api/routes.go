package main

import (
	"GoUber/pkg/handler"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")

	router.HandleFunc("/api/role", handler.GetAllRole).Methods("GET")
	router.HandleFunc("/api/role/{id:[0-9]+}", handler.GetRoleById).Methods("GET")

	return router
}
