package main

import (
	"flag"
	"log"

	"github.com/justyntemme/goNotifier/config"
	"github.com/justyntemme/goNotifier/web"
)

func main() {
	var configLocation string
	Config := new(config.Configuration)
	flag.StringVar(&configLocation, "config", "", "Config File location")
	flag.Parse()
	if configLocation == "" {
		log.Fatal("Please enter config file location")
	}
	config.ParseConfig(configLocation, Config)
	web.StartServer(Config.Port)
}
