// https://github.com/shichao-an/gopl.io/blob/master/ch1/echo1/main.go
// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func echo() {
	var s, sep string // implicitly initialized to the zero value for the type, which is "" for strings
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
