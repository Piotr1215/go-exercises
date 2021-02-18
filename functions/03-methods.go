package functions

import (
	"errors"
	"fmt"
)

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, "and I'm", p.age, "years old")
}

func EmployeeMe(name string, age int) (employee, error) {

	em, err := createEmployee(name, age)
	if err != nil {
		fmt.Println("Employee created", em)
	} else {
		fmt.Println("Errors", err)
	}

	return em, nil

}

func createEmployee(name string, age int) (employee, error) {

	e := employee{
		name,
		age,
	}
	if name == "" {
		return e, errors.New("Name cannot be empty")
	}

	if age < 0 {
		return e, errors.New("Age cannot be less than 0")
	}

	return e, nil
}
