package exercises

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

// Structs runs all functions in structs file
func Structs() {
	people := createSliceOfPerson()
	mapOfPeople := createMapFromPersons(people)

	fmt.Println(mapOfPeople)
}

func createMapFromPersons(people []person) map[string]person {

	mapOfPeople := map[string]person{
		people[0].lastName: people[0],
		people[1].lastName: people[1],
	}
	return mapOfPeople
}

func createSliceOfPerson() []person {
	// Initialization of struct
	people := []person{
		{
			firstName: "Atosik",
			lastName:  "Doggo",
			age:       7,
		},
		{
			firstName: "Piotr",
			lastName:  "Human",
			age:       100,
		},
	}
	return people
}
