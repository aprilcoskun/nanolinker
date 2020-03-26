package routes

import (
	"github.com/aprilcoskun/nanolinker/db"
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func logInPage(c *gin.Context) {
	if !db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/configure")
		return
	}
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func configurationPage(c *gin.Context) {
	if db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "configure.tmpl", nil)
}

func logIn(c *gin.Context) {
	if !db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/configure")
		return
	}

	var configData models.ConfigureUserData
	if err := c.ShouldBind(&configData); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := db.AuthUser(&configData)
	if err != nil {
		if err == db.ErrNotConfigured {
			c.String(http.StatusTemporaryRedirect, err.Error())
			return
		}

		logger.Error(err)
		c.String(http.StatusForbidden, err.Error())
		return
	}

	setSession(c, &configData)

	c.String(http.StatusOK, "user logged in")
}

func configuration(c *gin.Context) {
	var configData models.ConfigureUserData
	if err := c.ShouldBind(&configData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed!",
			"error":   err.Error(),
		})
		logger.Error(err.Error())
		return
	}
	err := db.Configure(&configData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		logger.Error(err)
		return
	}

	setSession(c, &configData)

	logger.Info("First Configuration Completed")
	c.String(http.StatusOK, "first configuration completed")
}
