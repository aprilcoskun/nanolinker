package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityMiddleWare(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(SecurityMiddleWare)

	r.GET("/test", func(c *gin.Context) {

		poweredBy := c.Writer.Header().Get("X-Powered-By")
		if poweredBy != "ZendServer" {
			c.Status(500)
			t.Error("X-Powered-By is corrupted:", poweredBy)
			return
		}

		frameOptions := c.Writer.Header().Get("X-Frame-Options")
		if frameOptions != "deny" {
			c.Status(500)
			t.Error("X-Frame-Options is corrupted:", frameOptions)
			return
		}

		contentTypeOptions := c.Writer.Header().Get("X-Content-Type-Options")
		if contentTypeOptions != "nosniff" {
			c.Status(500)
			t.Error("X-Content-Type-Options is corrupted:", contentTypeOptions)
			return
		}

		xssProtection := c.Writer.Header().Get("X-XSS-Protection")
		if xssProtection != "1; mode=block" {
			c.Status(500)
			t.Error("X-XSS-Protection is corrupted:", xssProtection)
			return
		}

		dnsPrefetchControl := c.Writer.Header().Get("X-DNS-Prefetch-Control")
		if dnsPrefetchControl != "off" {
			c.Status(500)
			t.Error("X-DNS-Prefetch-Control is corrupted:", dnsPrefetchControl)
			return
		}

		downloadOptions := c.Writer.Header().Get("X-Download-Options")
		if downloadOptions != "noopen" {
			c.Status(500)
			t.Error("X-Download-Options is corrupted:", downloadOptions)
			return
		}

		c.Status(200)
	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)

	r.ServeHTTP(resp, c.Request)
}
