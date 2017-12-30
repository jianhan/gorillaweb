package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type httpError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func newHTTPError(code uint, message string) *httpError {
	return &httpError{
		Code:    code,
		Message: message,
	}
}

func (h *httpError) Error() string {
	return fmt.Sprintf("[%d] %s", h.Code, h.Message)
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
