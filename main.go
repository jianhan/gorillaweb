package main

import (
	"github.com/davecgh/go-spew/spew"
	_ "github.com/jianhan/gorillaweb/bootstrap"
	"github.com/jianhan/gorillaweb/server"
	"github.com/spf13/viper"
)

func main() {
	spew.Dump(viper.Get("appName"))
	server.Run()
}
