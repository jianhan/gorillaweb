package auth

import (
	"errors"
	"net/http"
	"strings"

	auth0 "github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type jwtRequestValidatorScopeChecker interface {
	ValidateRequest(r *http.Request) error
	CheckScope(r *http.Request) error
}

type auth0ValidatorScopeChecker struct {
	domain       string
	clientID     string
	clientSecret string
	jwtValidator *auth0.JWTValidator
	token        *jwt.JSONWebToken
}

func NewJWTRequestValidatorScopeChecker(domain, clientID, clientSecret string, audiences []string) *auth0ValidatorScopeChecker {
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
	validator := auth0.NewValidator(auth0.NewConfiguration(client, audiences, apiIssuer, jose.RS256))
	return &auth0ValidatorScopeChecker{
		domain:       domain,
		clientID:     clientID,
		clientSecret: clientSecret,
		jwtValidator: validator,
	}
}

func (a *auth0ValidatorScopeChecker) ValidateRequest(r *http.Request) error {
	token, err := a.jwtValidator.ValidateRequest(r)
	if err != nil {
		return err
	}
	a.token = token
	return nil
}

func (a *auth0ValidatorScopeChecker) CheckScope(r *http.Request) error {
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
