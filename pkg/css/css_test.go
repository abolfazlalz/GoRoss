package css

import (
	"fmt"
	"testing"
)

func TestStyle_String(t *testing.T) {
	style := Style{
		Background: "red",
		Display:    Flex,
		Color:      "back",
	}
	fmt.Println(style)
}
