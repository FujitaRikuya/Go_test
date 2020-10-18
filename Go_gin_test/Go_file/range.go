package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 1}
	sum := 0

	for _, aa := range a {
		sum += aa
	}

	fmt.Println("range of a :", sum)
}
