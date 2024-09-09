package main

import "fmt"

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range n {
		if v%2 == 0 {
			fmt.Println("The number is even ", v)
		} else {
			fmt.Println("The number is odd", v)
		}
	}
}
