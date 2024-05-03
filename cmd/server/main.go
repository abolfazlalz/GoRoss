package main

import (
	"goross"
)

func main() {
	ui := goross.New()

	ui.Button("Hello World")
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
