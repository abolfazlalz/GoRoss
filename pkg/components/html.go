package components

type HTML struct {
	*Component
	title string
	meta  map[string]string
}

func NewHTML() *HTML {
	return &HTML{
		Component: &Component{},
		meta: map[string]string{
			"viewport": "width=device-width, initial-scale=1.0",
		},
	}
}

func (c *HTML) Title(title string) *HTML {
	c.title = title
	return c
}

func (c *HTML) AddMeta(key, value string) {
	c.meta[key] = value
}

func (c *HTML) GetMeta(key string) string {
	return c.meta[key]
}

func (c *HTML) DelMeta(key string) {
	delete(c.meta, key)
}

//  <meta name="viewport" content="width=device-width, initial-scale=1.0">

func (c *HTML) Render() string {
	head := &Component{TagName: "head"}
	html := &Component{TagName: "html"}

	title := &Component{TagName: "title", content: c.title}
	head.AddChild(title)

	for name, content := range c.meta {
		head.AddChild(&Component{TagName: "meta", attributes: ComponentProperty{"name": name, "content": content}})
	}

	html.AddChild(head)
	html.AddChild(c.children...)
	return html.Render()
}
