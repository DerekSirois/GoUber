package main

import (
	"GoUber/pkg/handler"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")

	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")

	router.HandleFunc("/api/role", handler.GetAllRole).Methods("GET")
	router.HandleFunc("/api/role/{id:[0-9]+}", handler.GetRoleById).Methods("GET")

	router.HandleFunc("/api/closest", handler.VerifyJWT(handler.GetClosestDriver)).Methods("GET")

	return router
}
