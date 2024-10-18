package main

import (
	"fmt"

	"github.com/daiqingyang/tools"
)

func main() {
	fmt.Println(tools.LineInFile([]byte("123445566 host"), "/etc/hosts", 0644))
}
