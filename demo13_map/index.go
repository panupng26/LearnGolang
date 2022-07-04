package main

import (
	"fmt"
)

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Println(numbers)

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#0ff"
	fmt.Println(colors)
	fmt.Println(colors["green"])

	var courses = make(map[string]map[string]int)

	// courses["Android"] = map[string]int{"price":200}
	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200
	courses["Android"]["code"] = 400

	courses["iOS"] = make(map[string]int)
	courses["iOS"]["price"] = 200
	courses["iOS"]["code"] = 400
	fmt.Println(courses["iOS"])
	fmt.Println(courses)
}