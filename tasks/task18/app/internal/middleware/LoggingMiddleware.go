package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		CallTime := time.Now()
		c.Next()
		method := c.Request.Method
		url := c.Request.URL.Path
		log.WithFields(logrus.Fields{
			"method": method,
			"url":    url,
			"Time":   CallTime,
		}).Info("Executed: ")
		c.Next()
	}
}
