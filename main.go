package main

import (
	"github.com/aprilcoskun/nanolinker/routes"
	"github.com/aprilcoskun/nanolinker/utils"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
)

func main() {
	// Use gin release mode on in production
	if strings.ToLower(os.Getenv("GO_ENV")) != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routes
	r := gin.New()

	r.Use(logger.Middleware, cors.Default(), utils.SecurityMiddleWare)
	r.LoadHTMLGlob("templates/*")
	routes.InitRoutes(r)

	logger.Info("Http Server Started at http://127.0.0.1:" + os.Getenv("PORT"))

	logger.Fatal(r.Run())
}
