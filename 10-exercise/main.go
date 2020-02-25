package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%b\t %d\t %x\n", num, num, num)
	num = num << 1
	fmt.Printf("%b\t %d\t %x\n", num, num, num)
}
