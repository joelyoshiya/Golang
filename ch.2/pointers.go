package main

import "fmt"

func main() {
	x := 1          // every variable has an address
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	// zero-value for a pointer is nil
}
