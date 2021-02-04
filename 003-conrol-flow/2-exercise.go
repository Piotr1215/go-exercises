package main

import "fmt"

func main() {
	for i := 33; i < 122; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t%#U\n", i, i)

		}
	}
}
