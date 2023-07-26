package main

import "fmt"

const LoginToken string = "valuegisd" // public- first letter capital

func main() {
	var username string = "shreyash"
	fmt.Println(username)
	fmt.Printf("variable is pf type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is pf type: %T \n", isLoggedIn)

	// uint uint8 uint16....

	// float32 float64

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "hijihij"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
