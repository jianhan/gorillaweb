package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var apiRoutes = Routes{
	Route{
		"Private",
		[]string{"GET"},
		"/private",
		private,
	},
}

func private(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}
