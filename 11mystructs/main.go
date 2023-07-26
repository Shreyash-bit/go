package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang, no super child parent

	shreyash := User{"Shreyash", "shreyash@bum.dev", true, 20}
	fmt.Println(shreyash)
	fmt.Printf("Shreyash details are: %v\n", shreyash)
	fmt.Printf("Name: %v\n", shreyash.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
