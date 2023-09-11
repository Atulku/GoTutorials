package main

import (
	"fmt"
	"math"
)

func main() {
	/*----***** Variable
	- go is case sensitive
	- no limit on length of var length
	- A un-used var or const will cause compiler error in go
	*/

	/*
		1. var  <var_name> <var_type>
				var a int              --> default value = 0
	*/

	var a int
	fmt.Println(a)

	/*
		2. var  <var_name> <var_type>
				var a int = 5          --> 5
	*/
	var b int = 5
	fmt.Println(b)
	/*
		3. multi var decl with no initial value
				var <name1> <name2> <var_type>
				var a, b int           --> 0, 0
	*/
	var c, d int
	fmt.Println(c, d)
	/*
		4. multi var decl with initial value
				var <name1> <name2>,....<name N> <var_type> = <val1>, <val2>, ....<valN>
				var a, b int = 8, 9    --> 8, 9
	*/
	var e, f int = 3, 4
	fmt.Println(e, f)
	/*
		5. Var with differnt type
				var (
					a int
					b int = 5
					c string = "abc"
				)
	*/
	var (
		g int    = 5
		h string = "multi var"
		i int
	)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	/*
		6. Type Inference: without specifying the type. Go compiler figures out the type based on the value assigned.
		type inference table:
			Integers				int
			Floats					float64
			Complex Numbers			complex128
			Strings					string
			Booleans				bool
			Characters				int32 or rune

		- For other types such as Array, Pointer, Structure etc, type inference will happen based on the value
	*/
	var j = 12
	var k = "circle"
	var l = 4.5
	var m = true
	var n = 'a'
	var o = 2 + 3i
	//var p = sample{name: "test"}          TODO

	fmt.Printf("Type: %T Value: %v\n", j, j)
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Printf("Type: %T Value: %v\n", l, l)
	fmt.Printf("Type: %T Value: %v\n", m, m)
	fmt.Printf("Type: %T Value: %v\n", n, n)
	fmt.Printf("Type: %T Value: %v\n", o, o)
	//fmt.Printf("Type: %T Value: %v\n", p, p)

	type sample struct {
		name string
	}
	/*
		7. short variable declaration using := operator
		when := operator is used both var and type can be omitted

		- := is only available within the function

		- If a variable i declared using :=, can't be redeclared using :=
			q := 7 --> error

	*/

	q := 5
	r := "string value"
	s := 5.7
	t := true

	fmt.Printf("Type: %T Value: %v\n", q, q)
	fmt.Printf("Type: %T Value: %v\n", r, r)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", t, t)

	// := can be used to declare multiple var in a single line
	u, v := 5, "hello"
	fmt.Println(u, v)

	// in multiple declaration can re-declare a var if one of the var in left side is new
	v, w := "world", 4.5
	fmt.Println(v, w)

	/*
		important points:
		- same var declared in the inner scope will shadow the var in outher scope

					package main
					import "fmt"
					var a = 123

					func main() {
						var a = 456
						fmt.Println(a)       --> 456
					}
		type of variable can't be changed later. Applicable for := as well

					var z int = 5
					z = "3.5"    // error

					z := 3
					z := "test"  // error
					z = 4  // valid
	*/

	// can assign value to var using expression as well

	x := 5 + 6
	y := math.Max(5, 7)
	fmt.Println(x, y)
}
