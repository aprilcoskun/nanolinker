package routes

import (
	"github.com/aprilcoskun/nanolinker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(router *gin.Engine) {
	router.StaticFile("/favicon.ico", "public/favicon.ico")
	// Init Session Store
	initSessionStore(router)

	// Init Web Routes
	router.GET("/", redirectHomePage)
	router.GET("/login", logInPage)
	router.GET("/configure", configurationPage)
	router.StaticFS("/public", http.Dir("public/"))
	router.GET("/l/:link", redirectLink)

	// Init Api Routes
	v1 := router.Group("/v1")

	// Guard Routes
	v1.Use(utils.AuthMiddleware)

	v1.GET("/", homePage)

	v1.POST("/login", logIn)
	v1.POST("/configure", configuration)

	v1.POST("/link", createLink)
	v1.PUT("/link", editLink)
	v1.DELETE("/link/:id", deleteLink)
}
