package main

import "fmt"

func main() {
	year := 1978
	i := 1
	for year <= 2021 {
		fmt.Println("I was", i, "in year", year)
		i++
		year++
	}
}
