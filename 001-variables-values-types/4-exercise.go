package main

import "fmt"

type age int

var x age

// Printf formats string, Println just prints value
func main() {
	fmt.Println(x)
	fmt.Printf("%T\t", x)

	x = 42
	fmt.Println(x)
}
