package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/v1adis1av28/level2/tasks/task18/app/internal/config"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/handler"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/models"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/server"
)

func main() {
	cfg := config.NewConfig()
	data := make(map[int][]models.Event, 256)
	handler := handler.NewHandler(data)
	server := server.NewServer(cfg, handler)

	go server.Start()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}
