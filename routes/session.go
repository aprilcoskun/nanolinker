package routes

import (
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/aprilcoskun/sessions"
	"github.com/aprilcoskun/sessions/memstore"
	"github.com/gin-gonic/gin"
	"os"
)

var sessionStore sessions.Store
var defaultSessionOptions sessions.Options

func initSessionStore(router *gin.Engine) {
	secret := os.Getenv("SESSION_KEY")
	defaultSessionOptions = sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		Path:     "/",
	}
	// If secret is not set in env variables, use default
	if secret == "" {
		secret = "nanolinker_secret"
	}

	sessionStore = memstore.NewStore([]byte(secret))
	router.Use(sessions.Sessions("nanolinker-api", sessionStore))
	return
}

func setSession(c *gin.Context, configData *models.ConfigureUserData) {
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
	}
}
