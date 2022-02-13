package domain

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func (p Person) ToStringMainData() string {
	val := fmt.Sprintf("%v, %v - ContactInfo: Email: %v, ZipCode: %v", p.LastName, p.FirstName, p.Email, p.ZipCode)
	return val
}

func (p Person) WrongUpdateName(newFirstName string) {
	p.FirstName = newFirstName
}
