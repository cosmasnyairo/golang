package main

import "fmt"

func main() {
	// map key value {}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["yellow"] = "sun"
	// colors["green"] = "grass"
	// colors["xanadu"] = "crosshair"

	colors := map[string]string{
		"yellow": "sun",
		"green":  "grass",
		"xanadu": "crosshair",
	}

	// delete(colors, "yellow")

	printMap(colors)

}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("color is", k, "object is", v)
	}
}
