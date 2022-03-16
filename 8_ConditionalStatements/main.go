package main

import "fmt"

func main() {
	fmt.Print("Enter your age: ")
	var age int16

	fmt.Scan(&age)

	if age > 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You are under 18")
	}
}
