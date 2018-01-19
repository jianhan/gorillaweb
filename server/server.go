package server

import (
	"log"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	gHandlers "github.com/jianhan/gorillaweb/handlers"
	"github.com/jianhan/gorillaweb/opts"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func Run(options opts.Options) {
	// connect to the database
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close() // clean up when we're done

	n := negroni.Classic()
	n.Use(gHandlers.MongoMiddleware(db))
	n.UseHandler(context.ClearHandler(gHandlers.AttachRouter(mux.NewRouter())))
	n.Run(":8888")
	// srv := &http.Server{
	// 	Handler:           n,
	// 	Addr:              fmt.Sprintf("%s:%d", viper.Get("server.host"), options.GetPort()),
	// 	WriteTimeout:      time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
	// 	ReadTimeout:       time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
	// 	ReadHeaderTimeout: time.Duration(viper.GetInt64("server.readHeaderTimeout")) * time.Second,
	// }
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
}
