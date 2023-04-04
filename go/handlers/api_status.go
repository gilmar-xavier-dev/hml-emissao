package handlers

import (
	"fmt"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetStatus: [%s]", r.Body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
