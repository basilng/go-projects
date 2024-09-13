package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Println(time.Since(start).Seconds())
//	fmt.Println(strings.Join(os.Args," ")) // e 1.1

}
