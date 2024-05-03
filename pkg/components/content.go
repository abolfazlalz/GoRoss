package components

type Content struct {
	ComponentI
}

func NewContent() *Content {
	return &Content{
		ComponentI: NewDiv(DivProps{
			classes: []string{"goross-content"},
		}),
	}
}
