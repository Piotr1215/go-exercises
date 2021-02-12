package main

import (
	"fmt"
)

func main() {

	x := [5]int{}

	for i := range x {
		x[i] = i
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x)
}
