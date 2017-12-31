package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
	_ "github.com/jianhan/gorillaweb/bootstrap"
	"github.com/jianhan/gorillaweb/server"
)

var opts struct {
	Port int `short:"p" long:"port" description:"port of app to run"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Port: %v\n", opts.Port)
	server.Run(server.NewServerOptions(opts.Port))
}
