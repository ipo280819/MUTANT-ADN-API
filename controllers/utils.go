package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseOK(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func ResponseStatus(status int, w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func ResponseError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %v", err)
}
func ResponseErrorStatus(status int, w http.ResponseWriter, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %v", err)
}
