package main

import "fmt"

func main() {
	fmt.Println("Class of pointers")

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23

	// a reference mean & ampersand
	var numPtr = &myNumber
	fmt.Println("address of pointer is ", numPtr)
	fmt.Println("Value of pointer is ", *numPtr)

	*numPtr = *numPtr + 2

	fmt.Println("address of actual pointer is ", numPtr)
	fmt.Println("New Value is:", myNumber)
}
