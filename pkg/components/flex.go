package components

import "goross/pkg/utils"

type Flex struct {
	ComponentI
}

type FlexDirection string

const (
	FlexRowDirection    FlexDirection = "row"
	FlexColumnDirection FlexDirection = "column"
)

func (fd FlexDirection) class() string {
	if fd == FlexRowDirection {
		return "flex-row goross-row"
	} else if fd == FlexColumnDirection {
		return "flex-column goross-column"
	}
	return ""
}

type FlexProps struct {
	DivProps
	Direction FlexDirection
	Wrap      bool
}

func NewFlex(propsOptional ...FlexProps) *Flex {
	props := utils.FirstSliceItem(propsOptional)
	props.AddClass("flex")
	if props.Direction != "" {
		props.AddClass(props.Direction.class())
	}
	if props.Wrap {
		props.AddClass("wrap")
	}

	return &Flex{
		NewDiv(props.DivProps),
	}
}
