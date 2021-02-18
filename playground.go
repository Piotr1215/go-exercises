package main

import (
	ex "github.com/Piotr1215/go-learning/basics"
	fun "github.com/Piotr1215/go-learning/functions"
	point "github.com/Piotr1215/go-learning/pointers"
)

func main() {

	// ex.Variables()
	// ex.ControlFlow()
	// ex.GroupingData()
	ex.Structs()
	fun.Functions()
	fun.Entry()
	fun.EntryAnonymous()
	fun.EntryFirstClassCitizens()
	fun.CallBackFunc()
	fun.CodeBlock()
	point.ExamplesOfPointers()
	point.EntryMethodSet()

	/* 		people := createSliceOfPerson()
	   		fmt.Println(people)

	   		for _, v := range people {
	   			fmt.Println(v.firstName, "is", v.age, "years old")
	   		} */

}
