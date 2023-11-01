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

	err = writeJson(w, payloadGenerator(false, "User created successfully"), http.StatusOK)
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

	err = writeJson(w, payloadGenerator(false, "You are logged in", token), http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}

func GetClosestDriver(w http.ResponseWriter, r *http.Request) {
	d, err := user.GetClosestDriverService(r)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJson(w, payloadGenerator(false, "Get closest driver", d), http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
}
