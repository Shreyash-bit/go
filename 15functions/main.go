package main

import "fmt"

func main() {
	fmt.Println("functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("result is: ", result)

	proRes := proAdder(3, 4, 5, 6, 7)
	fmt.Println("Pro Res: ", proRes)
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namastey")
}

func adder(a int, b int) int {
	return a + b
}
