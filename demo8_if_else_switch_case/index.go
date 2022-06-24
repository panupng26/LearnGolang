package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("value is 10")
	} else {
		fmt.Println("not value is 10")
	}

	if someValue >= 10 && someValue <= 20 {
		fmt.Println("this length 10-20")
	} else {
		fmt.Println("this not length 10-20")
	}
}

func fnSwitchCase() {
	index := 2
	switch index {
	case 0:
		fmt.Println("this 0")
	case 1:
		fmt.Println("this 1")
	}
}
