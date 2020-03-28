package routes

import (
	"github.com/aprilcoskun/nanolinker/db"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(c *gin.Context) {
	links, count, err := db.GetLinks(20, 0)
	if err != nil {
		logger.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "home.tmpl", &map[string]interface{}{
		"links":       links,
		"totalCount":  count,
		"linksLength": len(links),
	})
	return
}

func redirectHomePage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/v1")
}
