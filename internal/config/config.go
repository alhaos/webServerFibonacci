package config

import (
	"github.com/alhaos/webServerFibonacci/internal/webServer"
	"github.com/ilyakaznacheev/cleanenv"
)

// Configuration general app config
type Configuration struct {
	WebServer webServer.Config `yaml:"webServer"`
}

// New constructor for Configuration
func New(filename string) (*Configuration, error) {
	var c Configuration
	err := cleanenv.ReadConfig(filename, &c)
	if err != nil {
		return nil, err
	}
	return &c, err
}
