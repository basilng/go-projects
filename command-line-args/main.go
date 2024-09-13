package main

import (
	"fmt"
	"os"
)

// the below program is an example of unix echo command.
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "

	}
	fmt.Println(s)
}
