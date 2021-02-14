package exercises

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicleType vehicle
	fourWheel   bool
}

type sedan struct {
	vehicleType vehicle
	luxury      bool
}

// Structs runs all functions in structs file
func Structs() {
	people := createSliceOfPerson()
	mapOfPeople := createMapFromPersons(people)

	fmt.Println(mapOfPeople)
	cars()
	anonymStruct()
}

func anonymStruct() {
	repo := struct {
		url   string
		owner string
	}{
		url:   "https://google.com",
		owner: "Google",
	}

	fmt.Println(repo)
}

func cars() {
	truck := truck{
		vehicleType: vehicle{doors: 5, color: "red"},
		fourWheel:   true,
	}
	sedan := sedan{
		vehicleType: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(truck, sedan)
	fmt.Println(truck.vehicleType.color, sedan.luxury)
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
