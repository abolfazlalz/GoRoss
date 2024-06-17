package components

type HTML struct {
	*Component
	title string
	meta  map[string]string
}

func NewHTML() *HTML {
	return &HTML{
		Component: &Component{},
		title:     "GoRoss",
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

func (c *HTML) addCss(file string) {
	c.AddChild(&Component{TagName: "link", attributes: map[string]string{"rel": "stylesheet", "href": file}})
}

func (c *HTML) Render() string {
	head := NewHead()
	html := &Component{TagName: "html"}

	head.AddChild(NewScript().Src("/static/vue.global.prod.js"))

	head.Title(c.title)

	for name, content := range c.meta {
		head.AddChild(&Component{TagName: "meta", attributes: ComponentProperty{"name": name, "content": content}})
	}

	head.Stylesheet("/static/quasar.prod.css")
	head.Stylesheet("/static/goross.css")

	html.AddChild(head)

	body := &Component{TagName: "body"}

	appDiv := NewDiv(DivProps{id: "app"})

	layout := NewLayout()
	appDiv.AddChild(layout)

	content := NewContent()
	content.AddChild(c.children...)
	content.SetClass("goross-content")

	layout.AddChild(content)

	body.AddChild(appDiv)

	body.AddChild(NewScript().Src("/static/quasar.umd.prod.js"))
	body.AddChild(NewScript().Src("/static/tailwindcss.min.js"))
	body.AddChild(NewScript().Src("https://cdn.socket.io/socket.io-1.2.0.js"))
	body.AddChild(NewScript().Src("/static/goross.js"))
	html.AddChild(body)

	return html.Render()
}
