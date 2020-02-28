package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says Good morning, James.`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, shaken not stirred`)
}

func main() {
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa1)
}
