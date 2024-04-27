/**
align-items: normal;
align-items: stretch;

align-items: center;
align-items: start;
align-items: end;
align-items: flex-start;
align-items: flex-end;
align-items: self-start;
align-items: self-end;

align-items: baseline;
align-items: first baseline;
align-items: last baseline;
align-items: safe center;
align-items: unsafe center;

align-items: inherit;
align-items: initial;
align-items: revert;
align-items: revert-layer;
align-items: unset;
*/

package css

import (
	"fmt"
	"reflect"
)

type Display string

const (
	Flex  Display = "flex"
	Grid  Display = "grid"
	Block Display = "block"
)

type Alignment string

const (
	NormalAlignment    Alignment = "normal"
	StartAlignment     Alignment = "start"
	EndAlignment       Alignment = "end"
	FlexStartAlignment Alignment = "flex-start"
	FlexEndAlignment   Alignment = "flex-end"

	SelfStartAlignment Alignment = "self-start"
	SelfEndAlignment   Alignment = "self-end"

	BaselineAlignment  Alignment = "baseline"
	FirstBaselineAlign Alignment = "first baseline"
	LastBaselineAlign  Alignment = "last baseline"
	SafeCenterAlign    Alignment = "safe center"
	UnsafeCenterAlign  Alignment = "unsafe center"

	InheritAlignment      Alignment = "inherit"
	InitialAlignment      Alignment = "initial"
	RevertAlignment       Alignment = "revert"
	RevertLayoutAlignment Alignment = "revert-layer"
	UnsetAlignment        Alignment = "unset"
)

type Style struct {
	Background     string    `css:"background"`
	Color          string    `css:"color"`
	Display        Display   `css:"display"`
	AlignItems     Alignment `css:"align-items"`
	JustifyContent Alignment `css:"justify-content"`
}

func (s Style) String() string {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	str := ""

	// Iterate through the fields
	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Tag.Get("css")
		if name == "" {
			name = t.Field(i).Name
		}
		field := v.Field(i)
		value := field.String()
		if value == "" {
			continue
		}
		str += fmt.Sprintf("%s:%s;", name, field.String())
	}

	return str
}
