package db

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

type DB interface {
}

type MongoDB struct {
	mgo *mgo.Session
}

func NewDB() (DB, func(), error) {
	closeFunc := func() {}
	db, err := mgo.Dial(viper.GetString("mongo.url"))
	if err != nil {
		return nil, nil, err
	}
	closeFunc = func() {
		db.Close()
	}
	return &MongoDB{
		mgo: db,
	}, closeFunc, nil
}
