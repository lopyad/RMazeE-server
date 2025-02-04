package config

import (
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	if value := os.Getenv("PORT"); value == "" {
		c.Server.Port = "8080"
	} else {
		c.Server.Port = value
	}

	return c
}
