package main

import (
	_ "github.com/jianhan/gorillaweb/src/bootstrap"
	"github.com/jianhan/gorillaweb/src/opts"
	"github.com/jianhan/gorillaweb/src/server"
)

func main() {
	server.Run(opts.GetOpts())
}
