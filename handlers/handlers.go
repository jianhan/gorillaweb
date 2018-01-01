package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err = w.Write(body); err != nil {
		// TODO: log
	}
}

func AttachRouter(h *mux.Router) *mux.Router {
	r := newRoom()
	h.Handle("/chat", &templateHandler{filename: "chat.html"}).Name("chat")
	h.Handle("/room", r)
	go r.run()
	return h
}
