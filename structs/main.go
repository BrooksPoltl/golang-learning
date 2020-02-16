package main

import "fmt" 

type person struct {
	firstName string
	lastName string
	contactInfo contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	brooks := person{
		firstName: "Brooks",
		lastName: "Poltl",
		contactInfo: contactInfo{
			email: "bpoltl1@gmail.com",
			zipCode: 62234,
		},
	}
	brooks.updateName("tanner")
	brooks.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(name string) {
	p.firstName = name
}