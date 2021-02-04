package main

import "fmt"

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

// Formatting is pretty intuitive, very similar to F#
func main() {
	s := fmt.Sprintf(" Meaning of life: %v\n Name of a classy spy: %v\n Is Go awesome? %t", x, y, z)
	fmt.Println(s)
}
