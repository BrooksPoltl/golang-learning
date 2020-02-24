package main

import "fmt"

func main() {
	test := 4
	fmt.Println(test)
	fmt.Printf("%T\n", test)
	fmt.Printf("%b\n", test)
	fmt.Printf("%x\n", test)
	fmt.Printf("%#x\n", test)
	test = 911
	fmt.Printf("%#x\t%b\t%x\n", test, test, test)

	s := fmt.Sprintf("%#x\t%b\t%x\n", test, test, test)
	fmt.Println(s)
	fmt.Printf("%v\n", test)
}
