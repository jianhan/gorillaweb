package handlers

import (
	"context"
	"net/http"

	"github.com/jianhan/gorillaweb/src/auth"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func checkJWTMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	jvs := auth.NewJWTRequestValidatorScopeChecker(
		viper.GetString("auth0.domain"),
		viper.GetString("auth0.client_id"),
		viper.GetString("auth0.client_secret"),
		[]string{viper.GetString("auth0.audience")},
	)
	err := jvs.ValidateRequest(r)
	if err != nil {
		sendJSONResponse(rw, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
		return
	}
	err = jvs.CheckScope(r)
	if err != nil {
		sendJSONResponse(rw, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
		return
	}
	next(rw, r)
}

func MongoMiddleware(session *mgo.Session) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// copy the database session
		dbsession := session.Copy()
		defer dbsession.Close() // clean up
		ctx := r.Context()
		ctx = context.WithValue(ctx, "database", dbsession)
		r = r.WithContext(ctx)
		next(rw, r)
	})
}
