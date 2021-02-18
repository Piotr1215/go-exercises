package functions

import "fmt"

// CodeBlock shows how to scope a variable in a function inside a code block
func CodeBlock(){
	x := 100
	fmt.Println(x)
	// random code block
	{
		x := 200
		fmt.Println(x)
	}

}