package handler

import (
	"encoding/json"
	"net/http"
)

func PlaceholderGet(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "this is a placeholder GET response"})
}

func PlaceholderPost(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "this is a placeholder POST response"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
