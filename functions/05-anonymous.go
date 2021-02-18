package functions

import "fmt"

// EntryAnonymous showcases use of anonymous function
func EntryAnonymous(){
	name := "Atosik"

	func(){
		fmt.Println(name)
	}()

	func(name string){
		fmt.Println(name)
	}("Piotr")
}

