package config

import (
	"flag"
	"fmt"
	"time"
)

type Config struct {
	Host    string
	Port    int
	TimeOut time.Duration
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	timeout := flag.Int("timeout", 10, "timeout in seconds for connection")

	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		return nil, fmt.Errorf("usage: program [--timeout=10] host port")
	}

	cfg.Host = args[0]
	_, err := fmt.Sscanf(args[1], "%d", &cfg.Port)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %v", err)
	}
	if cfg.Port <= 0 || cfg.Port > 65535 {
		return nil, fmt.Errorf("port must be between 1 and 65535")
	}

	cfg.TimeOut = time.Duration(*timeout) * time.Second

	return cfg, nil
}
