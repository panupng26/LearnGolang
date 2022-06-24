package main

import "fmt"

func main() {
	// var numbers = make([]int, 0, 5) // 3 len ,5 cap

	var numbers []int
	showSlice(numbers)
}

func showSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}