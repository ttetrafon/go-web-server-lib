package api

import "net/http"

func RespondWithCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-missing-field", "myName")
	w.WriteHeader(http.StatusBadRequest)
}
