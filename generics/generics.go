package main

import "fmt"

func print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	print([]string{"Hello, ", "playground\n"})
	print([]int{1, 2, 3})
}
