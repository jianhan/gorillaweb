package opts

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
)

type Options interface {
	GetPort() int
}

type Args struct {
	Port int `short:"p" long:"port" description:"port of app to run"`
}

func (a Args) GetPort() int {
	if a.Port > 0 {
		return a.Port
	}
	return viper.GetInt("server.port")
}
func GetOpts() *Args {
	o := new(Args)
	_, err := flags.Parse(o)
	if err != nil {
		panic(err)
	}
	return o
}
