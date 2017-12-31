package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string
}

func examplePrivateHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, Person{"James"})
}

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

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func AttachRouter(h *mux.Router) *mux.Router {
	// jwtValidator := newJWTRequestValidatorScopeChecker(
	// 	viper.GetString("auth0.domain"),
	// 	viper.GetString("auth0.client_id"),
	// 	viper.GetString("auth0.client_secret"),
	// 	[]string{viper.GetString("auth0.audience")},
	// )
	r := newRoom()
	h.Handle("/room", r)
	// get the room going
	go r.run()
	// h.HandleFunc("/api/private", checkJWTHandler(examplePrivateHandler, jwtValidator))

	h.Handle("/chat", &templateHandler{filename: "chat.html"}).Name("home")
	// h.HandleFunc("/hello/{name}", index).Methods("GET")
	// h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("This is a catch-all route"))
	// }).Name("home")
	return h
}
