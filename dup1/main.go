package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) // . creating the input scanner object
	for input.Scan() {                  // scans in while loop until ctrl + d
		counts[input.Text()]++ // store the text to map key and incrementing the value count by 1
		/*
			the above statement is equal to the following
			line := input.Text()
			counts[line] = counts[line] + 1
			the expression evalueates from right to left for initially the counts[line] will be 0 since the line not exist in map
		*/
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
