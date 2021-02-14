package variables

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

// AssignValues assign
func AssignValues() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func PackageLevelScope() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func OutputFormatting() {
	s := fmt.Sprintf(" Meaning of life: %v\n Name of a classy spy: %v\n Is Go awesome? %t", x, y, z)
	fmt.Println(s)
}

// PrintVsPrintln Printf formats string, Println just prints value
func PrintVsPrintln() {
	fmt.Println(a)
	fmt.Printf("%T\t", a)

	a = 42
	fmt.Println(a)
}
