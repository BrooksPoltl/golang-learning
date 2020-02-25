package main

import "fmt"

const (
	a = 1 + iota
	b = 1 + iota
	c = 1 + iota
	d = 1 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
