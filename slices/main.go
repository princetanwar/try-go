package main

import "fmt"

func main() {
	var mySlice = []string{"ME"}

	mySlice = append(mySlice, "you")
	fmt.Println("my slice is ", mySlice)
	mySlice = append(mySlice[1:])

	fmt.Println("my slice is ", mySlice)
}
