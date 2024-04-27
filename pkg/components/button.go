package components

type Button struct {
	*Component
}

func NewButton() *Button {
	return &Button{
		&Component{TagName: "button"},
	}
}
