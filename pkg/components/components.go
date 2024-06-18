package components

import (
	"fmt"
	"goross/pkg/css"
	"goross/pkg/utils"
	"goross/pkg/websocket"
	"sort"
	"strings"
)

type ComponentProps struct {
	ID      string
	OnClick func()
}

type ComponentProperty map[string]string

func (cp ComponentProperty) String() string {
	text := ""
	for k, v := range cp {
		if v == "" {
			continue
		}
		if text != "" {
			text += " "
		}
		text += fmt.Sprintf("%s=\"%s\"", k, v)
	}
	if text == "" {
		return ""
	}
	return " " + text
}

type ComponentI interface {
	Render() string
	SetChildren(children ...ComponentI) ComponentI
	AddChild(child ...ComponentI) ComponentI
	SetClass(classNames ...string) ComponentI
	Classes() []string
	Style(style css.Style) ComponentI
	SetContent(content string) ComponentI
	InnerText() string
	InnerHtml() string
	Attr(name, value string) ComponentI
}

type Component struct {
	children   []ComponentI
	content    string
	TagName    string
	ID         string
	classes    []string
	style      css.Style
	attributes ComponentProperty
}

func NewComponent(tagName string, cProps ...ComponentProps) *Component {
	props := utils.FirstSliceItem(cProps)

	component := &Component{
		TagName: tagName,
		ID:      props.ID,
	}
	if props.OnClick != nil {
		ws := websocket.Instance()
		ws.Subscribe(func(s string) {
			if s == props.ID {
				props.OnClick()
			}
		})
		component.Attr("@click", fmt.Sprintf("handleClick('%s')", props.ID))
	}

	return component
}

func (c *Component) Render() string {
	childrenText := c.content
	for _, child := range c.children {
		childrenText += child.Render()
	}

	var props ComponentProperty
	if c.attributes != nil {
		props = c.attributes
	} else {
		props = ComponentProperty{}
	}
	props["class"] = strings.Join(c.classes, " ")
	props["style"] = c.style.String()
	props["id"] = c.ID

	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := ComponentProperty{}
	for _, k := range keys {
		result[k] = props[k]
	}

	return fmt.Sprintf("<%s%s>%s</%s>", c.TagName, props, childrenText, c.TagName)
}

func (c *Component) SetChildren(children ...ComponentI) ComponentI {
	c.children = children
	return c
}

func (c *Component) AddChild(child ...ComponentI) ComponentI {
	c.children = append(c.children, child...)
	return c
}

func (c *Component) SetClass(classNames ...string) ComponentI {
	c.classes = classNames
	return c
}

func (c *Component) Classes() []string {
	return c.classes
}

func (c *Component) Style(style css.Style) ComponentI {
	c.style = style
	return c
}

func (c *Component) SetContent(content string) ComponentI {
	c.content = content
	return c
}

func (c *Component) InnerText() string {
	return c.content
}

func (c *Component) InnerHtml() string {
	text := c.content
	for _, child := range c.children {
		text += child.Render()
	}
	return text
}

func (c *Component) Attr(name, value string) ComponentI {
	if c.attributes == nil {
		c.attributes = ComponentProperty{}
	}
	c.attributes[name] = value
	return c
}
