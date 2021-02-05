package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var prgName, firstLetter string

	if os.Args[0] != "" {
		prgName = os.Args[1:]
	} else {
		prgName = "Decoder"
	}

	firstLetter = string(prgName[0])

	if firstLetter == strings.ToUpper(firstLetter) {
		fmt.Println("Name of", prgName, "is correctly capitalized")
	} else {
		fmt.Println("Name of", prgName, "is not correctly capitalized")
	}
}
