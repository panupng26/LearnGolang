package main

import "fmt"

func main() {
	fmt.Println("Hello my name is bank")
	fnFor()
	fnWhile()

}

func fnFor()  {
	for index := 0; index < 10; index++ {
		fmt.Printf("Hello you are same : %d\n",index )
	}	
}

func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("Index %d\n", index)
		index++
	}
	
}

func fnWhileuseingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
	}
	
}

