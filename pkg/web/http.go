package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Web struct {
	*gin.Engine
}

//go:embed static/*
var contentFS embed.FS

func New() *Web {
	return &Web{
		Engine: assignRoutes(gin.New()),
	}
}

func NewWithFiber(app *gin.Engine) *Web {
	return &Web{
		Engine: assignRoutes(app),
	}
}

func assignRoutes(app *gin.Engine) *gin.Engine {
	app.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(contentFS))
	})

	return app
}

func (http Web) Run(addr ...string) error {
	return http.Engine.Run(addr...)
}
