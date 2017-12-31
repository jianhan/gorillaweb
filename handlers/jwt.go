package handlers

import "net/http"

func checkJWTHandler(
	handler func(w http.ResponseWriter, r *http.Request),
	jwtValidator jwtRequestValidatorScopeChecker,
) func(w http.ResponseWriter, r *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		err := jwtValidator.validateRequest(r)
		if err != nil {
			sendJSONResponse(w, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
			return
		} else {
			err := jwtValidator.checkScope(r)
			if err != nil {
				sendJSONResponse(w, http.StatusUnauthorized, newHTTPError(http.StatusUnauthorized, err.Error()))
				return
			} else {
				handler(w, r)
			}
		}

	}
	return h
}
