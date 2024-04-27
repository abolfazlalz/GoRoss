package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponent_AddChild(t *testing.T) {
	component := &Component{}
	child := &Component{}
	component.AddChild(child)
	assert.Equal(t, 1, len(component.children))
	assert.Equal(t, child, component.children[0])
}

func TestComponent_Classes(t *testing.T) {
	component := &Component{}
	classes := []string{"class1", "class2", "class3"}
	component.SetClass(classes...)
	assert.Equal(t, classes, component.Classes())
}
