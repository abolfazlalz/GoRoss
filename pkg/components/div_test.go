package components

import (
	"github.com/stretchr/testify/assert"
	"goross/pkg/css"
	"testing"
)

func TestNewDiv(t *testing.T) {
	div := NewDiv().AddChild(NewButton().SetClass("btn").SetContent("Click me")).AddChild(NewButton().SetContent("Hello World !")).Style(css.Style{Display: css.Flex, Background: "red"})
	expected := "<div style=\"background:red;display:flex;\"><button class=\"btn\">Click me</button><button>Hello World !</button></div>"

	assert.Equal(t, expected, div.Render())
}
