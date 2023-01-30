// Main allows testing of behavior of the package functions
package main

import (
	"fmt"
)

// make sure to "go build" instead of "go run" to use package level funcs

func main() {
	// command line args
	fmt.Println("Your command line args are:")
	os_args()
}
