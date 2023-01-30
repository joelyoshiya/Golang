package main

import "fmt"

func main() {
	// a full variable declaration
	var name string = "hello world"
	fmt.Println(name)

	// a short variable declaration
	name2 := "hello again"
	fmt.Println(name2)

	// := is declaration, while = is assignment
	// if := is re-used on a variable already declared, it behaves like assignment
	// you will typically only see this when assigning two or more values to variables
	// this will only work if introducing at least on new variables to the lhs

	// ex:
	// f, err := os.Open(infile)
	// ...
	// f, err := os.Create(outfile) // compile rror; no new variables

	// valid:
	// in, err := os.Open(infile)
	// ...
	// out, err := os.Create(outfile)

	// err will be updated via assignment, whereas out is declared
}
