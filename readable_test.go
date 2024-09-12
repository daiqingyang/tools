package tools

import (
	"fmt"
	"testing"
)

func TestHumanSize(t *testing.T) {
	fmt.Println(parseString("10GBx"))
	fmt.Println(parseString("10G"))
}

