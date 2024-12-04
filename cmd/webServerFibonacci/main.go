package main

import (
	"flag"
	"github.com/alhaos/webServerFibonacci/internal/config"
	"github.com/alhaos/webServerFibonacci/internal/webServer"
)

func main() {

	// Take file name
	confFilename := configFilename()

	// Init config
	conf, err := config.New(confFilename)
	if err != nil {
		panic(err)
	}

	// Init web server
	ws, err := webServer.New(conf.WebServer)
	if err != nil {
		panic(err)
	}

	// Run web server
	err = ws.Run()
	if err != nil {
		panic(err)
	}
}

// configFilename return config file from exec args
func configFilename() string {
	filenamePointer := flag.String("config", "config/config.yml", "app config filename")
	flag.Parse()
	filename := *filenamePointer
	return filename
}
