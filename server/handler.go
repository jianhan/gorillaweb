package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string
}

func examplePrivateHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, Person{"James"})
}
func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}
func AttachRouter(h *mux.Router) *mux.Router {
	jwtValidator := newJWTRequestValidatorScopeChecker(
		viper.GetString("auth0.domain"),
		viper.GetString("auth0.client_id"),
		viper.GetString("auth0.client_secret"),
		[]string{viper.GetString("auth0.audience")},
	)
	h.HandleFunc("/api/private", checkJWTHandler(examplePrivateHandler, jwtValidator))
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	}).Name("home")
	h.HandleFunc("/hello/{name}", index).Methods("GET")
	return h
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	spew.Dump(data, "DATA")
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
