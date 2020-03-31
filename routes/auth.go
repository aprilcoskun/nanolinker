package routes

import (
	"github.com/aprilcoskun/nanolinker/db"
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogInPage(c *gin.Context) {
	if !db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/configure")
		return
	}
	c.HTML(http.StatusOK, "login", nil)
}

func ConfigurationPage(c *gin.Context) {
	if db.IsConfigured() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "configure", nil)
}

func LogIn(c *gin.Context) {
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
	utils.SetSession(c, &configData)

	c.String(http.StatusOK, "user logged in")
}

func Configuration(c *gin.Context) {
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

	utils.SetSession(c, &configData)

	logger.Info("First Configuration Completed")
	c.String(http.StatusOK, "first Configuration completed")
}
