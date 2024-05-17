package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		c.Request.URL.Path = c.Param("proxyPath")

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
