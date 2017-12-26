package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	})
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", viper.Get("server.host"), viper.Get("server.port")),
		WriteTimeout: time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
		ReadTimeout:  time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
