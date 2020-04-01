package utils

import (
	"github.com/aprilcoskun/nanolinker/db"
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
	if !db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/configure/")
		c.Abort()
		return

	}
	// Check if user exist in session store
	if err != nil || statusCookie != "valid" || sessions.Default(c).Get("username") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login/")
		c.Abort()
		return
	}

	c.Next()
}
