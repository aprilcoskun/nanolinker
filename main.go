package main

import (
	"github.com/aprilcoskun/nanolinker/routes"
	"github.com/aprilcoskun/nanolinker/utils"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"os"
	"strings"
)

var templateBox = *packr.New("template", "templates/")
var publicBox = packr.New("public", "public/")

func main() {
	isDev := strings.ToLower(os.Getenv("GO_ENV")) == "dev"

	// Use gin release mode on in production
	if !isDev {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routes
	r := gin.New()

	// Use embedded Templates in Release Mode
	if isDev {
		r.LoadHTMLGlob("templates/*")
	} else {
		r.SetHTMLTemplate(mustLoadBoxedTemplate(templateBox))
	}

	r.Use(logger.Middleware, cors.Default(), utils.SecurityMiddleWare, utils.GzipMiddleware, utils.SessionMiddleware()).

		// Web Routes
		GET("/", routes.RedirectHomePage).
		GET("/login", routes.LogInPage).
		GET("/configure", routes.ConfigurationPage).
		GET("/l/:link", routes.RedirectLink).

		// Static files
		StaticFS("/public", publicBox).
		StaticFile("/favicon.ico", "public/favicon.ico")

	// Api Routes
	r.Group("/v1").
		Use(utils.AuthMiddleware).
		GET("/", routes.HomePage).
		POST("/login", routes.LogIn).
		POST("/configure", routes.Configuration).
		POST("/link", routes.CreateLink).
		PUT("/link/:id", routes.EditLink).
		DELETE("/link/:id", routes.DeleteLink)

	r.NoRoute(routes.RedirectLink)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("Http Server Started at http://127.0.0.1:" + port)
	logger.Fatal(r.Run())
}

// Embed Template files
// original source: https://github.com/gobuffalo/packr/issues/16#issuecomment-354905578
func mustLoadBoxedTemplate(box packr.Box) *template.Template {
	t := template.New("")
	err := box.Walk(func(path string, f packr.File) error {
		if path == "" {
			return nil
		}
		var err error
		var size int64
		if info, err := f.FileInfo(); err != nil {
			return err
		} else {
			size = info.Size()
		}

		// Normalize template name
		normalizedPath := path
		if strings.HasPrefix(path, "\\") || strings.HasPrefix(path, "/") {
			// don't want template name to start with / ie. /index.tmpl
			normalizedPath = normalizedPath[1:]
		}

		var h = make([]byte, 0, size)

		if h, err = box.Find(path); err != nil {
			return err
		}
		if _, err = t.New(normalizedPath).Parse(string(h)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic("error loading template")
	}
	return t
}
