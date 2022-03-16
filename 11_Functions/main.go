package main

import "fmt"

func main() {
	greetUsers("Rwitesh")

	fmt.Println(sum(6, 8))
	fmt.Println(getArea(5, 6))
}

func greetUsers(userName string) {
	fmt.Printf("Welcome, %v to our Conference!\n", userName)
}
