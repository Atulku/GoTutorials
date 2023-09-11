package main

import "fmt"

/*
	1. Local variable:
		-  block or function scope and lifetime
		-  Garbage collected after it's lifetime
	2. Global variable:
		- global within a package
		- if varaible name starts with lowercase, it can be accessed from the package within which it is defined.
		- if varaible name starts with uppercase, it can be accessed from the outside of the package.
		- lifetime of program
*/

var a = "global"

func main() {
	b := "local"

	fmt.Println(b)
	fmt.Println(a)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i) // error: undefined: i
}

func testScope() {
	fmt.Println(a)
	//fmt.Println(b) // error

}
