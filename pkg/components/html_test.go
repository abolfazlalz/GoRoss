package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHTML(t *testing.T) {
	html := NewHTML()
	html.Title("App 1")
	html.AddChild(NewButton().SetContent("Click me"))

	expected := "<html><head><title>App 1</title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"></meta></head><button>Click me</button></html>"
	assert.Equal(t, expected, html.Render())
}
