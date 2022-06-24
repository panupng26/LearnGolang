package main

import "fmt"

var count int = 0

func main() {
	fmt.Println("Begin");

	// Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true
	/*
	const tmp4 int = 0
	*/

	// Implicit Declaration
	tmp5 := 0
	tmp6 := "string bank"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp5)
	fmt.Println(tmp6)

	run()
}

func run()  {
	count++
	fmt.Println(count)
	count++
	fmt.Println(count)
	count++
	fmt.Println(count)
}