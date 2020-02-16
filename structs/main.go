package main

import "fmt" 

type person struct {
	firstName string
	lastName string
}

func main() {
	brooks := person{firstName: "Brooks", lastName: "Poltl"}
	fmt.Println(brooks)
}