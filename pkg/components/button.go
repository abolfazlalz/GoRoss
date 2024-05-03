package components

type Button struct {
	*Component
}

type ButtonProps struct {
	Color string
}

func (b ButtonProps) color() string {
	if b.Color == "" {
		return "primary"
	}
	return b.Color
}

func NewButton(content string, propsOptional ...ButtonProps) *Button {
	props := firstProp(propsOptional)
	btn := &Component{TagName: "q-btn"}
	btn.Attr("label", content)
	btn.Attr("color", props.color())
	return &Button{btn}
}
