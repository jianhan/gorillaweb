package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string
}

func examplePrivateHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, Person{"James"})
}

func checkJWTHandler(
	handler func(w http.ResponseWriter, r *http.Request),
	jwtValidator jwtRequestValidatorScopeChecker,
) func(w http.ResponseWriter, r *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		err := jwtValidator.validateRequest(r)
		if err != nil {
			sendJSONResponse(w, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
			return
		} else {
			err := jwtValidator.checkScope(r)
			if err != nil {
				sendJSONResponse(w, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
				return
			} else {
				handler(w, r)
			}
		}

	}
	return h
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
