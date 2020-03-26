package utils

import "github.com/gin-gonic/gin"

func SecurityMiddleWare(c *gin.Context) {
	// Security Headers
	c.Header("X-Powered-By", "ZendServer")
	c.Header("X-Frame-Options", "deny")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("X-DNS-Prefetch-Control", "off")
	c.Header("X-Download-Options", "noopen")
	c.Next()
}
