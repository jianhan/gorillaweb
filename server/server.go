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
	"github.com/spf13/viper"
)

func Run() {
	r := gHandlers.AttachRouter(mux.NewRouter())
	srv := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf("%s:%d", viper.Get("server.host"), viper.Get("server.port")),
		WriteTimeout:      time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
		ReadTimeout:       time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
		ReadHeaderTimeout: time.Duration(viper.GetInt64("server.readHeaderTimeout")) * time.Second,
	}
	if viper.GetBool("enableLog") {
		srv.Handler = handlers.LoggingHandler(os.Stdout, r)
	}
	log.Fatal(srv.ListenAndServe())
}
