package components

type Head struct {
	ComponentI
	title string
}

func NewHead() *Head {
	return &Head{
		ComponentI: &Component{TagName: "head"},
	}
}

func (h *Head) Title(title string) *Head {
	h.title = title
	return h
}

func (h *Head) Stylesheet(href string) *Head {
	link := NewLink()
	link.Rel("stylesheet")
	link.Href(href)
	h.AddChild(link)
	return h
}

func (h *Head) Render() string {
	title := &Component{
		TagName: "title",
		content: h.title,
	}
	h.AddChild(title)
	return h.ComponentI.Render()
}
