package main

import (
	"goross"
	"goross/pkg/components"
)

func main() {
	ui := goross.New()

	ui.Button("Hello World", components.ButtonProps{ID: "hello1"})
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
