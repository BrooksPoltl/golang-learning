package main

import "fmt"

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Printf("%T\t%T\n", a, b)
}
