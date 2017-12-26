package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AttachRouter(h *mux.Router) (*mux.Router, error) {
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	}).Name("home")
	return h, nil
}
