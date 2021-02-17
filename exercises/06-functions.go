package exercises

import "fmt"

type employee struct {
	Name string
	Age  int
}

type persona struct {
	first string
	last  string
	age   int
}

// Functions runs all funcs in Functions file
func Functions() {
	x, y := bar()
	fmt.Println(foo(), x, y)
	sumMe := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := sumInts(sumMe...)
	resSlice := sumIntsSlice(sumMe)

	fmt.Println(res, resSlice)

	deferExample()

	piotr := persona{
		first: "Piotr",
		last:  "Classified",
		age:   100,
	}

	piotr.speak()
}

func (p persona) speak() {
	fmt.Println("Hi, my name is", p.first, "and I'm", p.age, "years old")
}

func deferExample() {

	defer fmt.Println("I finish second although I appear first in the execution flow, the magic of defer!")
	fmt.Println("I finish first although I appear second in the execution flow, och defer what have u done!")
}

func (e employee) Init(name string, age int) employee {
	e.Name = name
	e.Age = age

	return e
}

func sumInts(ints ...int) int {
	s := 0
	for _, v := range ints {
		s += v
	}
	return s
}

func sumIntsSlice(ints []int) int {
	s := 0
	for _, v := range ints {
		s += v
	}
	return s
}

// create a func with the identifier foo that returns an int
func foo() int {
	return 100
}

// create a func with the identifier bar that returns an int and a string
func bar() (int, string) {
	return 100, "years"
}
