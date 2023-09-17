package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("a + b =>", add(1, 2))
	fmt.Println("a + b =>", sum(1, 2))
}
