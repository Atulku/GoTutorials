Go adv:
	- Simple syntax, easy to write code.
	- statically typed
	- inbuilt concurrency.
	- compiled language with fast compilation.
	- garbage collected
	- standard library and tooling for test, code formatting, linting, memory analysis
Go disadv: 
	- No inheritance in go. go does not have clases. Go favours composition over inheritance.
	- not a pure OO language
	- No generics in go. Need to write explicit code for what you want.
	- No compile time polymorphism and function overloading
--------------------------------------------------------------------------------------------
*******************------------GFG Go lang---------------******************

*** Identifiers
	- begin with letter or underscore
	- valid: a-z, A-z, 0-9, _
	- case sensitive
	- can not use keyword
	- no limit on length, but use optimal lenght 4-15.
	- Blank identifier:
			- underscore character(_) is known as a . 
	  		- It is used as an anonymous placeholder instead of a regular identifier.
	- Expoted Identifiers: 
			- can be accessed from other package
			rule: 
				- first character should be uppercase
				- Id should be declared in the package block., or it is a var name or fun name
*** Keywords
	- Total 25 keywords are present in go lang

*** Data Type
	-   1. Basic type: Numbers, strings, and booleans come under this category.
		2. Aggregate type: Array and structs come under this category.
		3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
		4. Interface type	

	** basic DT:
	   1. Numbers: 
	       - int8, int16, int32, int64
		   - uint8, uint16, uint32, uint64
		   - int, uint
		   - rune: it is synonym of int32, also represent unicode code points
		   - byte: synonym of uint8
		   - uintptr: width is not defined, but it cn hold all the bits of a ptr value

		2. Float:
		   - float32
		   - float64
		   	e.g:  a:= 45.40, var x float32 = 5.00

		3.  Complex Numbers:
		   - complex64      var a complex64 = complex(6, 2), var exp = 12+45i , exp := 23+34i
		   - complex128     var a complex128 = complex(6, 2)

		   - Built-in functions:
		   		a. comp := 13 + 34i

			        realNum : = real(comp)  --> 13
					imaginary := imag(comp) --> 34

		4. Boolean:
		   - represents only one bit of information.
		   - Value of boolean can not be converted to any other type(explicit or implicit)
		    
		      e.g:    result := 2 == 2
			  
		5. Strings:
		   - A sequence of immutable bytes. 
		   - Once string is created, can not be changed.
		   - Can use + operator for concatenation.
		      
		      e.g:    str := "Hello"
			          str1 := "World"

					  fmt.Println("New String :", str+str1)
**** Variable
    - go is case sensitive
    - no limit on length of var length.
	- Use Identifiers for variable
    
		1. var  <var_name> <var_type>
				var a int              --> default value = 0
		2. var  <var_name> <var_type>
				var a int = 5          --> 5
		3. multi var decl with no initial value 
				var <name1> <name2> <var_type>
				var a, b int           --> 0, 0
		4. multi var decl with initial value 
				var <name1> <name2>,....<name N> <var_type> = <val1>, <val2>, ....<valN>
				var a, b int = 8, 9    --> 8, 9
		5. Var with differnt type
				var (
					a int
					b int = 5
					c string = "abc"
				)
		6. Type Inference: without specifying the type. Go compiler figures out the type based on the value assigned.

		e.g: 
				1. var a,b,c int = 3,4,5
				2. var c,d,e = 2, 34.32, "hello"  // type is not used, can declare different types
				3. var file,err = os.Open(name)  // Can initialize a set of variable using a function which returns multiple values.
				4. num := 5	   // short variable declaration
				5. myvar1, myvar2, myvar3 := 800, 34, 56

*** Scope of variable	
    - in Golang all identifiers are lexically (or statically) scoped, i.e. scope of a variable can be determined at compile time.
	- 2 type:
		a. Local Variables(Declared Inside a block or a function): function/block lifetime
		b. Global Variables(Declared outside a block or a function): program lifetime
	- There will be a compile-time error if these variables are declared twice with the same name in the same scope.
    - What will happen there exists a local variable with the same name as that of the global variable inside a function?
		ans: compiler will give preference to local variable
				
*** Constants
	- once the value of constant is defined, it cannot be modified further.
	- There can be any basic data type of constants like an integer constant, a floating constant, a character constant, or a string literal.
			a. Numeric Constant (Integer constant, Floating constant, Complex constant)
			b. String literals
			c. Boolean Constant	
	- Can not declare const by using := operator

				e.g: const str = "Hello"
					 const PI = 3.14
					 const correct = true
					 // typed and untyped const
					 const typedConst float64 = 123.12
					 const typedConstInt int = 123

					 const untypedInt = 23
					 const untypedFloat = 45.34
					 // Multiple const declaration
					 const (
						GFG     = "GeeksforGeeks"
						Correct = true
						Pi      = 3.14
					)
*** Access Specifier/ Escape sequence character:

    - \n : next line   
	- %d : Print Numbers
	- %f : Print Float
	- %s : Print String
	- %T : type of variable

*** Operators 
	- Arithmetic Operators : +, -, *, /, %
	- Relational Operators : ==, !=, <, <=, >, >= 
	- Logical Operators : &&, |, !
	- Bitwise Operators : &, |, ^, <<, >>, &^(AND NOT)
	- Assignment Operators : =, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
	

	- &^ (AND NOT): This is a bit clear operator. The &^ operator is bit clear (AND NOT): in the expression z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; 
	                   otherwise it equals the corresponding bit of x.

    - Misc Operators: 
		& : Address of variable
		* : pointer to a variable
		<- : Receive operator, is used to receive a value from the channel
*** Rune: ToDo

*** Type casting













// Queries:
- How a declare char.
- " " and ' ' for string declaration
- &^ (AND NOT): This is a bit clear operator. The &^ operator is bit clear (AND NOT): in the expression z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the corresponding bit of x.
- rune
- <- Receive operator