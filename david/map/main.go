package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// no dot notion in maps, have to use square braces. Keys have to be same data type and values have to be same data typte. 3 different ways to create maps.
//1st way
// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#4bf745",
// }

//2nd way
// var colors map[string]string

//3rd way
// colors := make(map[string]string)

//to add values

// colors["white"] = "#ffffff"

//to delete keys

// delete(colors, "white")
