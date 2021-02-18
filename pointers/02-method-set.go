package pointers

import "fmt"

type person struct {
	first string
	last  string
}

func EntryMethodSet() {
	p1 := person{
		"Piotr",
		"Clasified",
	}

	fmt.Println(p1)
	fmt.Println(&p1)


	changeMe(&p1)

	fmt.Println(p1)
	fmt.Println(&p1)
}

func changeMe(p *person) {
	(*p).first = "Atosik"
	(*p).last = "Psiaczek"

	fmt.Println(&p)
}