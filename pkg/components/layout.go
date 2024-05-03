package components

type Layout struct {
	ComponentI
}

func NewLayout(propsOptional ...DivProps) *Layout {
	props := firstProp(propsOptional)

	props.classes = append(props.classes, "layout")

	div := NewDiv(props)
	div.TagName = "q-layout"

	return &Layout{
		div,
	}
}
