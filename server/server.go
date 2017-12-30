package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
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

type jwtRequestValidatorScopeChecker interface {
	validateRequest(r *http.Request) error
	checkScope(r *http.Request) error
}

type auth0ValidatorScopeChecker struct {
	domain       string
	clientID     string
	clientSecret string
	jwtValidator *auth0.JWTValidator
	token        *jwt.JSONWebToken
}

func newJWTRequestValidatorScopeChecker(domain, clientID, clientSecret string, audiences []string) *auth0ValidatorScopeChecker {
	// start validation for constructor
	if strings.TrimSpace(domain) == "" {
		panic("Domain can not be empty")
	}
	if strings.TrimSpace(clientID) == "" {
		panic("Client ID can not be empty")
	}
	if strings.TrimSpace(clientSecret) == "" {
		panic("Client secret can not be empty")
	}
	if len(audiences) == 0 {
		panic("Audiences can not be empty")
	}
	// start build struct
	jwksURI := "https://" + domain + "/.well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI})
	apiIssuer := "https://" + domain + "/"
	configuration := auth0.NewConfiguration(client, audiences, apiIssuer, jose.RS256)
	validator := auth0.NewValidator(configuration)
	return &auth0ValidatorScopeChecker{
		domain:       domain,
		clientID:     clientID,
		clientSecret: clientSecret,
		jwtValidator: validator,
	}
}

func (a *auth0ValidatorScopeChecker) validateRequest(r *http.Request) error {
	token, err := a.jwtValidator.ValidateRequest(r)
	if err != nil {
		return err
	}
	a.token = token
	return nil
}

func (a *auth0ValidatorScopeChecker) checkScope(r *http.Request) error {
	claims := map[string]interface{}{}
	if a.jwtValidator == nil {
		return errors.New("jwtValidator is nil")
	}
	if a.token == nil {
		return errors.New("token is nil, please validate request first, then check scope")
	}
	err := a.jwtValidator.Claims(r, a.token, &claims)
	if err != nil {
		return err
	}
	// TODO: scope not setup just yet
	// if claims["scope"] != nil && strings.Contains(claims["scope"].(string), "read:messages") {
	// 	return nil
	// }
	if claims["scope"] != nil {
		return nil
	}
	return errors.New("Invalid scope")
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
