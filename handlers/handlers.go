package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}
func AttachRouter(h *mux.Router, avs *auth0ValidatorScopeChecker) *mux.Router {
	r := newRoom()
	h.Handle("/chat", &templateHandler{filename: "chat.html"}).Name("chat")
	h.Handle("/room", r)
	h.HandleFunc("/private/{name}", checkJWTHandler(index, avs)).Methods("GET").Name("private")
	go r.run()
	return h
}
