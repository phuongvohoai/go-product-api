package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var memoryCache = cache.New(3*time.Second, 6*time.Second)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func ResponseCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		key := c.Request.URL.String()

		if getCachedResponse(c, key) {
			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		if c.Writer.Status() == http.StatusOK {
			setCachedResponse(key, w.body.Bytes())
		}

	}
}

func getCachedResponse(c *gin.Context, key string) bool {
	if cachedResponse, found := memoryCache.Get(key); found {
		log.Printf("[DEBUG] Cache hit for key: %s", key)
		if str, ok := cachedResponse.(string); ok {
			c.Writer.WriteString(str)
			c.Abort()
			return true
		} else {
			c.AbortWithStatusJSON(200, cachedResponse)
			return true
		}
	}
	return false
}

func setCachedResponse(key string, body []byte) {
	var responseObject map[string]interface{}
	err := json.Unmarshal(body, &responseObject)
	if err != nil {
		log.Print("[ERROR] Failed to unmarshal response body: ", err)
		memoryCache.Set(key, string(body), cache.DefaultExpiration)
		return
	} else {
		memoryCache.Set(key, &responseObject, cache.DefaultExpiration)
	}
}
