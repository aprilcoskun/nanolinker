package routes

import (
	"github.com/aprilcoskun/nanolinker/db"
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func notFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "notfound.tmpl", nil)
}

func redirectLink(c *gin.Context) {
	linkId := c.Param("link")
	if linkId == "" {
		linkId = c.Request.URL.String()[1:]
	}

	link, err := db.GetLink(linkId)
	if err != nil {
		notFoundPage(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link.Url)
	err = db.InsertClick(&models.Click{
		LinkID:    link.ID,
		Ip:        c.ClientIP(),
		Referer:   c.Request.Referer(),
		UserAgent: c.Request.UserAgent(),
	})
	if err != nil {
		logger.Error(err)
	}
}

func createLink(c *gin.Context) {
	var cachedLink models.CachedLink
	if err := c.ShouldBind(&cachedLink); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := db.SaveLink(cachedLink)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "link saved")
}

func deleteLink(c *gin.Context) {
	err := db.DeleteLink(c.Param("id"))
	if err != nil {
		logger.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "link deleted")
}

func editLink(c *gin.Context) {
	var cachedLink models.CachedLink
	if err := c.ShouldBind(&cachedLink); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := db.UpdateLink(c.Param("id"), cachedLink)
	if err != nil {
		logger.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "link updated")
}
