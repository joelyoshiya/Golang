package main

import (
	"fmt"
	"os"
)

func os_args() {
	fmt.Println(os.Args[1:]) // args is a slice of strings, prints all args except the first (program name/comand)
}
