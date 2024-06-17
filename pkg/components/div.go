package components

import "goross/pkg/utils"

type Div struct {
	*Component
}

type DivProps struct {
	classes []string
	id      string
}

func (d *DivProps) AddClass(class ...string) {
	if d.classes == nil {
		d.classes = make([]string, 0)
	}
	d.classes = append(d.classes, class...)
}

func NewDiv(propsOptional ...DivProps) *Button {
	props := utils.FirstSliceItem(propsOptional)

	return &Button{
		&Component{TagName: "div", classes: props.classes, ID: props.id},
	}
}
