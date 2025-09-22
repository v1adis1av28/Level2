package models

import "time"

type Config struct {
	BaseDir      string
	MaxDepth     int
	Url          string
	WorkersCount int
	Timeout      time.Duration
}

type Job struct {
	URL   string
	Depth int
}
