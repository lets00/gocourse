package main

import (
	"fmt"
)

var z = 10

func main() {
	// Dynamic attribution (this works on codeblock only)
	x := 10
	y := "Olá"
	fmt.Println(x, y, z)
	x = 20
	fmt.Println(x, y, z)
}
