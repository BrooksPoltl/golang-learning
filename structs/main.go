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
	fmt.Printf("%+v", brooks)
}