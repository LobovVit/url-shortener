package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	HostA string `env:"SERVER_ADDRESS"`
	HostB string `env:"BASE_URL"`
}

var Instance *Config

func GetConfig() *Config {
	Instance = &Config{}
	err := env.Parse(Instance)
	if err != nil {
		log.Fatal(err)
	}
	hostA := flag.String(`a`, "localhost:8080", "SERVER_ADDRESS")
	hostB := flag.String(`b`, "http://localhost:8080", "BASE_URL")
	flag.Parse()

	if Instance.HostA == "" {
		Instance.HostA = *hostA
	}
	if Instance.HostB == "" {
		Instance.HostB = *hostB
	}

	return Instance
}
