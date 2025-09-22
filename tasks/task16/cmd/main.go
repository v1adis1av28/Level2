package main

import (
	"time"

	"github.com/v1adis1av28/level2/tasks/task16/internal/app"
	"github.com/v1adis1av28/level2/tasks/task16/internal/models"
)

func main() {
	cfg := models.Config{
		MaxDepth:     2,
		WorkersCount: 10,
		Timeout:      30 * time.Second,
	}
	app.StartApp(&cfg)
}
