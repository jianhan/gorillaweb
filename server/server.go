package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gHandlers "github.com/jianhan/gorillaweb/handlers"
	"github.com/jianhan/gorillaweb/opts"
	"github.com/spf13/viper"
)

func Run(options opts.Options) {
	r := gHandlers.AttachRouter(mux.NewRouter())
	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("%s:%d", viper.Get("server.host"), options.GetPort()),
		WriteTimeout:      time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
		ReadTimeout:       time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
		ReadHeaderTimeout: time.Duration(viper.GetInt64("server.readHeaderTimeout")) * time.Second,
	}
	if viper.GetBool("enableLog") {
		srv.Handler = handlers.LoggingHandler(os.Stdout, r)
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
