package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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

func AttachRouter(h *mux.Router) *mux.Router {
	// define all api routes here
	apiRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	apiRouter.HandleFunc("/private", index).Methods("GET")
	// connect main route with middleware via negroni
	h.PathPrefix("/api/v1").Handler(negroni.New(
		negroni.HandlerFunc(checkJWTMiddleware),
		negroni.Wrap(apiRouter),
	))
	return h
}
