package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("a + b =>", add(1, 2))
}
