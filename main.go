package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jianhan/gorillaweb/bootstrap"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8888", loggedRouter)
}
