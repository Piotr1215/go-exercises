package functions

import (
	"fmt"
	"log"
)

func EntryFirstClassCitizens() {
	nameFunc := nameFunction
	nameFunc()

	returnFunc := returnFunction()

	fmt.Println(returnFunc())

}

func CallBackFunc(){
	// Functional composition
	addWithLogger(printer, 1,2,3,4,5)
	addWithLogger(logger, 1,2,3,4,5)
}

func printer(logme ...interface{}){
	fmt.Println(logme...)
}


func logger(logme ...interface{}){
	log.Println(logme...)
}

func addWithLogger(log func(x ...interface{}), numbers ...int){
	sum  := 0
	for _, v := range numbers {
		sum += v
		log("Adding", v)
	}

	log("Total sum is", sum)
}


func nameFunction() {
	fmt.Println("Hallo from nameFunction")
}

func returnFunction() func() string {
	return func () string {
		return "Atosik"
	}
}