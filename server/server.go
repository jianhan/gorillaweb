package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	gHandlers "github.com/jianhan/gorillaweb/handlers"
	"github.com/jianhan/gorillaweb/opts"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func Run(options opts.Options) {
	n := negroni.Classic()
	n.UseHandler(gHandlers.AttachRouter(mux.NewRouter()))
	srv := &http.Server{
		Handler:           n,
		Addr:              fmt.Sprintf("%s:%d", viper.Get("server.host"), options.GetPort()),
		WriteTimeout:      time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
		ReadTimeout:       time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
		ReadHeaderTimeout: time.Duration(viper.GetInt64("server.readHeaderTimeout")) * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
