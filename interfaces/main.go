package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hi There!"
}
func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
