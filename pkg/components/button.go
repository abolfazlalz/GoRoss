package components

import (
	"fmt"
	"goross/pkg/utils"
)

type Button struct {
	*Component
}

type ButtonProps struct {
	ComponentProps
	Color string
	ID    string
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

	btn.Attr("@click", fmt.Sprintf("handleClick('%s')", props.ID))
	btnC := &Button{btn}
	btnC.ID = props.ID
	return btnC
}
