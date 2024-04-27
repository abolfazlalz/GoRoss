package components

import (
	"github.com/stretchr/testify/assert"
	"goross/pkg/css"
	"testing"
)

func TestComponent_Render(t *testing.T) {
	btn := NewButton().Style(css.Style{Color: "red"}).SetContent("Click me")
	assert.Equal(t, "<button style=\"color:red;\">Click me</button>", btn.Render())
}
