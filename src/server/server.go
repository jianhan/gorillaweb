package server

import (
	"fmt"
	"log"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	gHandlers "github.com/jianhan/gorillaweb/src/handlers"
	"github.com/jianhan/gorillaweb/src/opts"
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
	n.UseHandler(context.ClearHandler(gHandlers.InitRoutes(mux.NewRouter())))
	n.Run(fmt.Sprintf(":%d", options.GetPort()))
}
