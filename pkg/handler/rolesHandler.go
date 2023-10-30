package handler

import (
	"GoUber/pkg/roles"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllRole(w http.ResponseWriter, _ *http.Request) {
	role, err := roles.GetAll()
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJson(w, payloadGenerator(false, "Get all roles", role), http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	role, err := roles.GetById(int(id))
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJson(w, payloadGenerator(false, "Get one roles", role), http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}
