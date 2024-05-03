package components

type Script struct {
	ComponentI
}

func NewScript() *Script {
	return &Script{
		ComponentI: &Component{TagName: "script"},
	}
}

func (s *Script) Src(src string) *Script {
	s.Attr("src", src)
	s.Attr("type", "text/javascript")
	return s
}
