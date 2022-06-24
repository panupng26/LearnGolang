package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	
	fmt.Println(msg)

	changeMessage(&msg,"change message")
	fmt.Println(msg)

	changeMessage(msgPointer,"Hello bank")
	fmt.Println(msg)

	changeMessage(msgPointer,"Hello bank 1")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string)  {
	*aPointer = newMessage
}