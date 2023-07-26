package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list :", fruitList)
	fmt.Println("Fruit list :", len(fruitList))

	var vegList = [6]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list:", len(vegList))
}
