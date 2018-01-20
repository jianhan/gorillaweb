package main

import (
	"fmt"
	"log"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jianhan/gorillaweb/api/bootstrap"
	"github.com/jianhan/gorillaweb/api/db"
	gHandlers "github.com/jianhan/gorillaweb/api/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	if err := bootstrap.Boot(); err != nil {
		panic(fmt.Sprintf("Bootstrap error: %s", err.Error()))
	}
	db.NewDB()
	// connect to the database
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close() // clean up when we're done
	n := negroni.Classic()
	n.Use(gHandlers.MongoMiddleware(db))
	n.UseHandler(context.ClearHandler(gHandlers.InitRoutes(mux.NewRouter())))
	logrus.Info("successfully started")
	n.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}
