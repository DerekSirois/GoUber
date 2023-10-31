package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool
	Message string
	Data    any    `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

func Index(w http.ResponseWriter, _ *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Welcome to GoUber",
	}

	err := writeJson(w, payload, http.StatusOK)
	if err != nil {
		ErrorJson(w, err, http.StatusInternalServerError)
	}
}

func readJson(r *http.Request, data any) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("body must have only a single JSON value")
	}
	return nil
}

func writeJson(w http.ResponseWriter, data any, status int, headers ...http.Header) error {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func ErrorJson(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	err = writeJson(w, payload, statusCode)
	if err != nil {
		log.Panic(err)
	}
}

func payloadGenerator(err bool, message string, data ...any) jsonResponse {
	var payload jsonResponse
	payload.Error = err
	payload.Message = message

	if len(data) > 0 {
		payload.Data = data[0]
	}

	return payload
}
