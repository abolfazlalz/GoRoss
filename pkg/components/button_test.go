package components

import (
	"github.com/stretchr/testify/assert"
	"goross/pkg/css"
	"testing"
)

func TestComponent_Render(t *testing.T) {
	btn := NewButton("Click me").Style(css.Style{Color: "red"})
	assert.Equal(t, "<button style=\"color:red;\">Click me</button>", btn.Render())
}
