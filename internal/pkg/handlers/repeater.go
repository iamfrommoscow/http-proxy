package handlers

import (
	"net/http"
	"proxy/internal/pkg/db"
	"strconv"

	"github.com/gorilla/mux"
)

func RepeatRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	request, err := db.SelectRequest(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var client http.Client
	resp, err := client.Do(request)
}
