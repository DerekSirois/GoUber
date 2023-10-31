package handler

import (
	"GoUber/pkg/user"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	u := user.User{}
	err := readJson(r, &u)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	err = user.CreateService(u)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "User created successfully",
	}

	err = writeJson(w, payload, http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	u := user.Login{}
	err := readJson(r, &u)
	if err != nil {
		ErrorJson(w, err)
		return
	}

	uDb, err := user.GetByEmailDB(u.Email)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	if !user.CheckPasswordHash(u.Password, uDb.Password) {
		ErrorJson(w, fmt.Errorf("wrong email or password"), http.StatusUnauthorized)
		return
	}

	token, err := user.CreateJWTToken(uDb.Id, uDb.Email)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "You are logged in",
		Token:   token,
	}

	err = writeJson(w, payload, http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}
