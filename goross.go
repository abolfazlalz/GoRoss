package goross

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goross/pkg/components"
	"goross/pkg/utils"
	"goross/pkg/web"
	"net/http"
)

type GoRoss struct {
	container components.ComponentI
	id        int
}

func New() *GoRoss {
	html := components.NewHTML()

	return &GoRoss{
		container: html,
	}
}

func (ui *GoRoss) Button(content string, props ...components.ButtonProps) *GoRoss {
	c := utils.FirstSliceItem(props)
	c.ID = fmt.Sprintf("button_%d", ui.id)

	ui.id++
	ui.container.AddChild(components.NewButton(content, c))
	return ui
}

func (ui *GoRoss) Row(callback func(ui *GoRoss)) *GoRoss {
	old := ui.container
	defer func() {
		ui.container = old
	}()

	row := components.NewFlex(components.FlexProps{Direction: components.FlexRowDirection})

	ui.container.AddChild(row)
	ui.container = row

	callback(ui)

	return ui
}

func (ui *GoRoss) Run(addr ...string) error {
	req := web.New()

	req.GET("/", func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.String(http.StatusOK, ui.container.Render())
	})

	return req.Run(addr...)
}
