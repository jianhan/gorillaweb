package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err = w.Write(body); err != nil {
		log.Printf("Failed to write the response body: %v", err)
	}
}

func AttachRouter(h *mux.Router) *mux.Router {
	r := newRoom()
	h.Handle("/chat", &templateHandler{filename: "chat.html"}).Name("home")
	h.Handle("/room", r)
	// get the room going
	go r.run()
	return h
}
