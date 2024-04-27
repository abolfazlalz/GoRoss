package components

type Div struct {
	*Component
}

func NewDiv() *Button {
	return &Button{
		&Component{TagName: "div"},
	}
}
