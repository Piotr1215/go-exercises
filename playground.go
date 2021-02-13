package main

import (
	"fmt"
)

func main() {
	fmt.Println("Make examples")

	x := makePlay()
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1215)
	fmt.Println(x)
	fmt.Println(len(x))

	fmt.Println("Map examples")

	m := mapPlay()

	fmt.Println(m)

	fmt.Println(m[1])

	// Use v, ok (comma, ok) check to see if value is in a map

	if v, ok := m[1]; ok {
		fmt.Println("Value with key", 1, " is in the map and the value is... treble...", v)
	} else {
		fmt.Println("Value with key", 1, " is NOT in the map buuuuuu")
	}

}

func makePlay() []int {
	// Use make if you know how many elements u will have in an array
	x := make([]int, 10, 100)

	x[0] = 100
	x[9] = 999

	return x
}
func mapPlay() map[int]string {

	mapa := map[int]string{

		1: "Atosik",
		2: "Costam",
		3: "Learning Go",
	}

	return mapa
}
