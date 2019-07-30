package main

import (
	"fmt"
)

func main() {
	_, err := fmt.Println("Hello world!", 100)
	fmt.Println(err)
}
