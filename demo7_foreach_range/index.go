package main

import "fmt"

func main() {
	fmt.Println("bank")
	courses := []string {"Andorid","ios","react"}
	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1,item)
	}

	for _, item := range courses {
		fmt.Printf("this %s\n", item)
	}
}