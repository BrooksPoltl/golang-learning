package main

import "fmt" 

type person struct {
	firstName string
	lastName string
}

func main() {
	var brooks person
	brooks.firstName = "Brooks"
	brooks.lastName = "Poltl"
	fmt.Println(brooks)
}