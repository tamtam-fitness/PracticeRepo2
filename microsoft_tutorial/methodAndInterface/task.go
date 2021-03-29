package methodAndInterface

import "fmt"

type Account struct {
	firstName string
	lastName  string
}

func (a *Account) ChangeName(firstName string, lastName string) {
	a.firstName = firstName
	a.lastName = lastName
}

func (a *Account) String() string {
	return fmt.Sprintf("first: %v last: %v", a.firstName, a.lastName)
}

type Employee struct {
	Account
	credit float64
}

func (e *Employee) AddCredits(credit float64) {
	e.credit += credit
}

func (e *Employee) RemoveCredits(credit float64) {
	e.credit -= credit
}

func (e *Employee) CheckCredits() float64 {
	return e.credit
}

// ==============================================
