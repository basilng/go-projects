package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// 1. one way of declearing and initialzing map.

	// var colors map[string]string
	// 0 value map
	// 2. another way of using map without initializing

	// colors := make(map[string]string)
	//3. another way of declaring a map

	colors["white"] = "#ffffff"
	// accessing the value of the map

	// fmt.Println(colors)

	// delete(colors, "white")
	// deleting the key value from map.

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

/*
Printing map
*/
