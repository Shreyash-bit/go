package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang, no super child parent

	shreyash := User{"Shreyash", "shreyash@bum.dev", true, 20}
	fmt.Println(shreyash)
	fmt.Printf("Shreyash details are: %v.\n", shreyash)
	fmt.Printf("Name: %v and email is %v.\n", shreyash.Name, shreyash.Email)
	shreyash.GetStatus()
	shreyash.NewMail()
	fmt.Printf("Name: %v and email is %v.\n", shreyash.Name, shreyash.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
