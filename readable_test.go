package tools

import (
	"fmt"
	"testing"
)

func TestHumanSize(t *testing.T) {
	fmt.Println(HumanSize(16751259648))
	fmt.Println(parseString("10G"))
}

