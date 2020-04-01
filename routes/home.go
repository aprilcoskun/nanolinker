package routes

import (
	"github.com/aprilcoskun/nanolinker/db"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func HomePage(c *gin.Context) {
	pageCount := 0
	limit := 20
	offset := 0
	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err == nil {
		offset = (pageNumber - 1) * limit
	} else {
		pageNumber = 1
	}
	links, count, err := db.GetLinks(limit, offset)
	if err != nil {
		logger.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if count > limit {
		pageCount = int(math.Ceil(float64(count) / float64(limit)))
	}

	c.HTML(http.StatusOK, "home", &gin.H{
		"links":       links,
		"totalCount":  count,
		"linksLength": len(links),
		"pageCount":   pageCount,
		"pageNumber":  pageNumber,
	})
	return
}

func RedirectHomePage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/v1")
}
