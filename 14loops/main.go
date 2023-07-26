package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"Sunday", "Teusday", "Wednesday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 2 {
			goto lco
		}

		if roughValue == 5 {
			roughValue++
			continue
		}
		fmt.Println("value is:", roughValue)
		roughValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.com")

}
