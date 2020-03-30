package utils

import (
	"github.com/aprilcoskun/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	url := c.Request.URL.String()

	// Bypass auth routes
	if url == "/v1/login" || url == "/v1/configure" {
		c.Next()
		return
	}

	statusCookie, err := c.Cookie("session-status")
	// Check if user exist in session store
	if err != nil || statusCookie != "valid" || sessions.Default(c).Get("username") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login/")
		return
	}

	c.Next()
}
