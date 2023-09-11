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
	


----***** Variable
    - go is case sensitive
    - no limit on length of var length
    
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
			