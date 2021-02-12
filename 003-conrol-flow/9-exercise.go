package main

import (
	"fmt"
)

func main() {
	var favSport string = "Judo"

	switch favSport {
	case "Karate":
		fmt.Println("Favourite sport is Karate")
	case "Judo":
		fmt.Println("Favourite sport is judo")
	default:
		fmt.Println("No favourite sport")
	}
}
