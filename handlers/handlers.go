package handlers

import (
	"encoding/json"
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

func AttachRouter(h *mux.Router) *mux.Router {
	apiRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	for _, route := range apiRoutes {
		apiRouter.Methods(route.Method...).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	h.PathPrefix("/api/v1").Handler(negroni.New(
		negroni.HandlerFunc(checkJWTMiddleware),
		negroni.Wrap(apiRouter),
	))
	return h
}
