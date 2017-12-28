package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type auth0Configs struct {
	domain       string
	clientID     string
	clientSecret string
	audience     string
}

func (a *auth0Configs) newAuth0Configs(domain, clientID, clientSecret, audience string) *auth0Configs {
	return &auth0Configs{
		domain:       domain,
		clientID:     clientID,
		clientSecret: clientSecret,
		audience:     audience,
	}
}

func (a *auth0Configs) validate() error {
	if strings.TrimSpace(a.domain) == "" {
	}
}

func init() {

}

func Run() {
	r := AttachRouter(mux.NewRouter())

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
