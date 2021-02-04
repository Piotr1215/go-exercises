package main

import "fmt"

type age int

var x age
var y int

// Printf formats string, Println just prints value
func main(){
	fmt.Println(x)
	fmt.Printf("%T\t", x)

	x = 42

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\t", y)
	fmt.Println(x)
}