package main

import (
	"goross"
	"goross/pkg/components"
	"log"
)

func main() {
	ui := goross.New()

	ui.Button("Im clickable", components.ButtonProps{
		ComponentProps: components.ComponentProps{OnClick: func() {
			log.Println("im clickable !")
		}},
	})
	ui.Button("Click me")

	ui.Row(func(ui *goross.GoRoss) {
		ui.Button("Hello World")
		ui.Button("Click me")
	})

	ui.Button("Hello World")
	ui.Button("Click me")

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
