package components

type Link struct {
	ComponentI
}

func NewLink() *Link {
	return &Link{
		ComponentI: &Component{TagName: "link"},
	}
}

func (l *Link) Rel(rel string) *Link {
	l.Attr("rel", rel)
	return l
}

func (l *Link) Href(href string) *Link {
	l.Attr("href", href)
	return l
}
