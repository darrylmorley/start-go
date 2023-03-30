package main

import "fmt"

// A map is a collection of key/value pairs like an Object in Javascript

// Keys must all be the same type, values must all be the same type, keys and values can be different types

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fff",
	}

	// var colors map[string]string // use this if you want to assign values later on

	// colors := make(map[string]string) // Another option for a map to be used later

	// colors["white"] = "#fff" // Assign a key/vlue to a map

	// delete(colors, "white") // delete a key/value

	printMap(colors)
}

// Iterate a map
func printMap(c map[string]string) {
	for color, hex := range c {

		fmt.Println(color, hex)
	}
}
