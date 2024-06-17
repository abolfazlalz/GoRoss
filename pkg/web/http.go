package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"goross/pkg/websocket"
	"net/http"
)

type Web struct {
	*gin.Engine
	ws *websocket.WebSocket
}

//go:embed static/*
var contentFS embed.FS

func New() *Web {
	return &Web{
		Engine: assignRoutes(gin.New()),
		ws:     websocket.New(),
	}
}

func assignRoutes(app *gin.Engine) *gin.Engine {
	app.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(contentFS))
	})

	return app
}

func (http Web) Run(addr ...string) error {
	http.ws.Init()
	defer http.ws.Close()
	http.Engine.GET("/socket.io/*any", gin.WrapH(http.ws.Server()))
	http.Engine.POST("/socket.io/*any", gin.WrapH(http.ws.Server()))
	return http.Engine.Run(addr...)
}
