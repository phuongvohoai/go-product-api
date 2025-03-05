package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func PerformanceLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		switch {
		case latency < 100*time.Millisecond:
			log.Print("[DEBUG] Request latency: ", latency)
		case latency < 300*time.Millisecond:
			log.Print("[INFO] Request latency: ", latency)
		case latency > 300*time.Millisecond:
			log.Print("[WARN] Request latency: ", latency)
		}
	}
}
