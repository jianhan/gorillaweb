package handlers

import (
	"fmt"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var apiRoutes = Routes{
	Route{
		"Private",
		[]string{"GET"},
		"/private",
		private,
	},
}

type comment struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	Text   string        `json:"text" bson:"text"`
	When   time.Time     `json:"when" bson:"when"`
}

func private(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db := r.Context().Value("database").(*mgo.Session)
	// decode the request body
	c := comment{Author: "Jian", Text: "Description"}
	// give the comment a unique ID
	c.ID = bson.NewObjectId()
	c.When = time.Now()

	// insert it into the database
	if err := db.DB("commentsapp").C("comments").Insert(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Hello:")
}
