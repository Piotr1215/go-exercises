package main

import (
	"fmt"
)

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

type age int

var a age

func main() {
	assignValues()
	packageLevelScope()
	outputFormatting()
	printVsPrintln()
}

func assignValues() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func packageLevelScope() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func outputFormatting() {
	s := fmt.Sprintf(" Meaning of life: %v\n Name of a classy spy: %v\n Is Go awesome? %t", x, y, z)
	fmt.Println(s)
}

// Printf formats string, Println just prints value
func printVsPrintln() {
	fmt.Println(a)
	fmt.Printf("%T\t", a)

	a = 42
	fmt.Println(a)
}
