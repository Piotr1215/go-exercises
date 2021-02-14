package structsFun

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func createMapFromPersons()  {
	people := createSliceOfPerson()
	mapOfPeople := map[string]person{
		people[0].lastName: people[0],
		people[1].lastName: people[1],
	}
	fmt.Println(mapOfPeople)
}

func createSliceOfPerson() []person {
	// Initialization of struct
	people := []person{
		person{
			firstName: "Atosik",
			lastName:  "Doggo",
			age:       7,
		},
		person{
			firstName: "Piotr",
			lastName:  "Human",
			age:       100,
		},
	}
	return people
}
