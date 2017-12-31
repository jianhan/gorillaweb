package main

import (
	_ "github.com/jianhan/gorillaweb/bootstrap"
	"github.com/jianhan/gorillaweb/opts"
	"github.com/jianhan/gorillaweb/server"
)

func main() {
	server.Run(opts.GetOpts())
}
