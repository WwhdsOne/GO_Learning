package main

import "fmt"

type Account struct {
	firstName string
	lastName  string
}

func (a *Account) ChangeName(newName string) { // pointer receiver
	a.firstName = newName
}

type Employee struct {
	a       Account
	Credits float64
}

func (e *Employee) AddCredits(credits float64) {
	e.Credits = e.Credits + credits
}

func (e *Employee) RemoveCredits(credits float64) {
	e.Credits = e.Credits - credits
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func (a Account) String() string {
	return fmt.Sprintf("FirstName: %s, LastName: %s", a.firstName, a.lastName)
}

func main() {
	account := Account{firstName: "John", lastName: "Doe"}
	e := Employee{a: account, Credits: 100}
	fmt.Printf("%s\n", account.String())
	fmt.Println(account)
	fmt.Println(e)
	e.AddCredits(50)
	fmt.Println(e.CheckCredits())
	e.RemoveCredits(20)
	fmt.Println(e.CheckCredits())
}
