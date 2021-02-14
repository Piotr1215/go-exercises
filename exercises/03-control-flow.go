package exercises

import (
	"fmt"
	"strings"

)

// ControlFlow executes all control flow functions
func ControlFlow() {
	simpleLoop()
	nestedLoops()
	howOldInYear()
	howOldInYearWithBreak()
	printModuloFour()
	capitalizerCheck("Atosik")
	simpleSwitch()
	favoriteSportSwitch("Judo")
	logicDefaults()
}

// SimpleLoop that's all
func simpleLoop() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}
func nestedLoops() {
	for i := 33; i < 122; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t%#U\n", i, i)

		}
	}
}
func howOldInYear() {
	year := 1978
	i := 1
	for year <= 2021 {
		fmt.Println("I was", i, "in year", year)
		i++
		year++
	}
}
func howOldInYearWithBreak() {
	year := 1978
	i := 1
	for {
		if year > 2021 {
			break
		}
		fmt.Println("I was", i, "in year", year)
		i++
		year++
	}
}
func printModuloFour() {
	for i := 10; i < 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
func capitalizerCheck(name string) {

	firstLetter := string(name[0])

	if firstLetter == strings.ToUpper(firstLetter) {
		fmt.Println("Name of", name, "is correctly capitalized")
	} else {
		fmt.Println("Name of", name, "is not correctly capitalized")
	}
}
func simpleSwitch() {
	x := 10
	y := 10

	switch x == y {
	case true:
		fmt.Println("This is true")
	}
}
func favoriteSportSwitch(favSport string) {

	switch favSport {
	case "Karate":
		fmt.Println("Favourite sport is Karate")
	case "Judo":
		fmt.Println("Favourite sport is judo")
	default:
		fmt.Println("No favourite sport")
	}
}
func logicDefaults() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
