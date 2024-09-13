package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
	// here the _ is called blank identifier, go does not allow unused local variable
		s += sep + arg
		sep = " "
	}
	fmt.Println(s, " ", time.Since(start).Seconds())


//	for i, arg := range os.Args[1:] { // e 1.2
//		fmt.Println(i, " ", arg) 
//	}
}
