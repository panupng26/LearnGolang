package main

import "fmt"

func main() {
	var array1 []string = []string{"bank", "may"}
	var array2 []string = []string{"bank", "may", "god", "father"}
	for index, data := range array1 {
		fmt.Printf("%d = %s\n", index, data)
	}
	fmt.Println("")
	for index, data := range array2 {
		fmt.Printf("%d = %s\n", index, data)
	}
}