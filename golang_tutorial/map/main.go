package main

import "fmt"

func main() {
	// colors := make(map[string]int)
	// colors["a"] = 11
	// colors["b"] = 12
	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#54546",
	// }

	// fmt.Println(colors)
	// delete(colors, "a")
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#54546",
		"white": "#ffff",
	}
	printColors(colors)
	fmt.Println(colors)

}

func printColors(c map[string]string) {
	for color, _ := range c {
		fmt.Println("Hex code for", color, "is")
	}
	c["11"] = "123"
}
