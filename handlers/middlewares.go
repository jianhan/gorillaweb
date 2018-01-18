package handlers

import (
	"net/http"

	"github.com/spf13/viper"
)

func checkJWTMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	jvs := NewJWTRequestValidatorScopeChecker(
		viper.GetString("auth0.domain"),
		viper.GetString("auth0.client_id"),
		viper.GetString("auth0.client_secret"),
		[]string{viper.GetString("auth0.audience")},
	)
	err := jvs.validateRequest(r)
	if err != nil {
		sendJSONResponse(rw, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
		return
	}
	err = jvs.checkScope(r)
	if err != nil {
		sendJSONResponse(rw, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
		return
	}
	next(rw, r)
}
