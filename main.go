package main

import (
	"github.com/aprilcoskun/nanolinker/routes"
	"github.com/aprilcoskun/nanolinker/utils"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nanmu42/gzip"
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

	gzipMiddleware := gzip.NewHandler(gzip.Config{
		// gzip compression level to use
		CompressionLevel: 6,
		// minimum content length to trigger gzip, 8kb
		MinContentLength: 1024 * 8,
	}).Gin

	r.Use(logger.Middleware, cors.Default(), utils.SecurityMiddleWare, gzipMiddleware)
	r.LoadHTMLGlob("templates/*")
	routes.InitRoutes(r)

	logger.Info("Http Server Started at http://127.0.0.1:" + os.Getenv("PORT"))

	logger.Fatal(r.Run())
}
