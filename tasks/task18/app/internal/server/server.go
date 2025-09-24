package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/config"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/handler"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/middleware"
)

type Server struct {
	Config  *config.Config
	Router  *gin.Engine
	Handler *handler.Handler
}

func NewServer(cfg *config.Config, handler *handler.Handler) *Server {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	server := &Server{
		Config:  cfg,
		Router:  router,
		Handler: handler,
	}

	server.SetupRoutes()
	return server
}

func (s *Server) Start() {
	s.Router.Run(fmt.Sprintf(":%s", s.Config.Port))
}

func (s *Server) SetupRoutes() {
	s.Router.Use(middleware.LoggingMiddleware())
	s.Router.GET("/api/events_for_day", s.Handler.GetDayEvents)
	s.Router.GET("/api/events_for_week", s.Handler.GetWeekEvents)
	s.Router.GET("/api/events_for_month", s.Handler.GetMonthEvents)

	s.Router.POST("/api/create_event", s.Handler.CreateEvent)
	s.Router.POST("/api/update_event", s.Handler.UpdateEvent)
	s.Router.POST("/api/delete_event", s.Handler.DeleteEvent)
}
