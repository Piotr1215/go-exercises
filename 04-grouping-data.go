package main

import (
	"fmt"
)

var mapa = map[int]string{

	1: "Atosik",
	2: "Costam",
	3: "Learning Go",
}

func main() {
	intSlice()
	slicesOfSlice()
	appendToSlice()
	manipulateSlice()
	printStates()
	makeExample()
	mapExample()
	mapCommaOkCheck()
}
func mapCommaOkCheck() {
	// Use v, ok (comma, ok) check to see if value is in a map

	if v, ok := mapa[1]; ok {
		fmt.Println("Value with key", 1, " is in the map and the value is... treble...", v)
	} else {
		fmt.Println("Value with key", 1, " is NOT in the map buuuuuu")
	}
}
func mapExample() {

	fmt.Println(mapa)
}
func makeExample() {
	// Use make if you know how many elements u will have in an array
	x := make([]int, 10, 100)

	x[0] = 100
	x[9] = 999

	fmt.Println(x)
}

func intSlice() {

	x := [5]int{}

	for i := range x {
		x[i] = i
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x)
}

func slicesOfSlice() {
	x := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func appendToSlice() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 52, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	// append, "..." nad prepend
	x = append(x, y...)

	fmt.Println(x)
}
func manipulateSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}
func printStates() {
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(states)
}
