package main

import "fmt"

func main() {
	// var numbers = make([]int, 0, 100) // 3 len ,5 cap

	// var numbers []int
	numbers := []int{}
	// var numbers [5]int
	// fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 100,200)
	showSlice(numbers)
}

func showSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}