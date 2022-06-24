package main

import "fmt"

func main() {
	fn1()
	fn2("Good Morning")
	fn3("Good Morning", 1)
	const a,b = 2,3
	fmt.Printf("%d+%d => %d\n",a,b, sum(a,b))

	var x,y int = swap(1,0)
	fmt.Printf("%d+%d => %d\n",x,y, sum(x,y))

	z1,z2 := swapV2(a,b)
	fmt.Printf("%d+%d => %d\n",z1,z2, sum(z1,z2))
}

func fn1()  {
	fmt.Println("Bank panupong")
}

func fn2(msg string)  {
	fmt.Println(msg)
}

func fn3(title string, version int)  {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a int,b int) int  {
	return a+b
}

func swap(a int,b int) (int, int)  {
	return b,a
}

func swapV2(a,b int) (x,y int)  {
	y = a
	x = b
	return
}