package main

import "fmt"

func main() {
	// var myPointer *int

	// fmt.Println("my pointer value is", myPointer)

	var num = 23

	var myPtr = &num

	fmt.Println("pointer is", myPtr)
	fmt.Println("pointer value is ", *myPtr)

	*myPtr = *myPtr + 2
	fmt.Println("new pointer value is ", num, *myPtr)

}
