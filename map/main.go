package main

import "fmt"

func main() {
	// 1. Declare a map
	// colors := make(map[string]string)

	// 2. Declare a map
	// var colors map[string]string

	// 3. Declare a map
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4b7456",
	}

	// Assign map's property
	// colors["white"] = "#ffffff"
	
	// Delete map's property
	// delete(colors, "white")	

	printMap(colors)
}

func printMap (c map[string]string){
	// Loop through a map
	for color, hex := range c {
		fmt.Println("Hex for color ", color, " is ", hex)
	}
}