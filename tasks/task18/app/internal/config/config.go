package config

import "flag"

type Config struct {
	Port string
}

func NewConfig() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.Port, "port", "8080", "Port that server started on")

	flag.Parse()

	return cfg
}
