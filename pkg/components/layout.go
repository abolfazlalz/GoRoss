package components

import "goross/pkg/utils"

type Layout struct {
	ComponentI
}

func NewLayout(propsOptional ...DivProps) *Layout {
	props := utils.FirstSliceItem(propsOptional)

	props.classes = append(props.classes, "layout")

	div := NewDiv(props)
	div.TagName = "q-layout"

	return &Layout{
		div,
	}
}
