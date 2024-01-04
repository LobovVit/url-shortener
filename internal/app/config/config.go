package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	HostA string `env:"ADDRESS_A"`
	HostB string `env:"ADDRESS_B"`
}

var instance *Config

func GetConfig() *Config {
	instance = &Config{}
	err := env.Parse(instance)
	if err != nil {
		log.Fatal(err)
	}
	hostA := flag.String("a", "localhost:8080", "адрес эндпоинта HTTP-сервера")
	hostB := flag.String("a", "localhost:8080", "адрес эндпоинта HTTP-сервера")
	flag.Parse()

	if instance.HostA == "" {
		instance.HostA = *hostA
	}
	if instance.HostB == "" {
		instance.HostB = *hostB
	}
	return instance
}
