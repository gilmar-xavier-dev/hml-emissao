package handlers

import (
	"fmt"
	"net/http"
)

func CreateQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateQuery: [%s]", r.Body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CreateWithdrawalQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateWithdrawalQuery: [%s]", r.Body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
