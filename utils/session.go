package utils

import (
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/aprilcoskun/sessions"
	"github.com/aprilcoskun/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

//var sessionStore cookie.Store
var defaultSessionOptions = sessions.Options{
	MaxAge:   60 * 60 * 24 * 7,
	HttpOnly: true,
	Path:     "/",
}

func SessionMiddleware() gin.HandlerFunc {
	sessionKey := os.Getenv("SESSION_KEY")

	// If secret is not set in env variables, use default
	if sessionKey == "" {
		sessionKey = "nanolinker_secret"
	}
	return sessions.Sessions("nanolinker-api", cookie.NewStore([]byte(sessionKey)))
}

func SetSession(c *gin.Context, configData *models.ConfigureUserData) {
	session := sessions.Default(c)
	session.Set("username", configData.Username)
	sessionOptions := defaultSessionOptions

	// Save session for 10 years if remember me is enabled
	if configData.RememberMe {
		sessionOptions.MaxAge = 60 * 60 * 24 * 365 * 10
	}

	session.Options(sessionOptions)

	err := session.Save()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c.SetCookie(
		"session-status",
		"valid",
		sessionOptions.MaxAge,
		sessionOptions.Path,
		sessionOptions.Domain,
		sessionOptions.Secure,
		false)
}
