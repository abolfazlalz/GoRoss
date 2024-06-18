package components

import (
	"goross/pkg/utils"
)

type Button struct {
	*Component
}

type ButtonProps struct {
	ComponentProps
	Color string
}

func (b ButtonProps) color() string {
	if b.Color == "" {
		return "primary"
	}
	return b.Color
}

func NewButton(content string, propsOptional ...ButtonProps) *Button {
	props := utils.FirstSliceItem(propsOptional)
	btn := NewComponent("q-btn", props.ComponentProps)
	btn.Attr("label", content)
	btn.Attr("color", props.color())

	btnC := &Button{btn}
	btnC.ID = props.ID
	return btnC
}
