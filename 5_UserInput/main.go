package main

import "fmt"

func main() {
	var userName string
	var userPhone int64

	// Asking for user input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&userName) // & is a pointer

	fmt.Print("Enter your contanct number: ")
	fmt.Scan(&userPhone)

	fmt.Println("Username :", userName)
	fmt.Println("Contact :", userPhone)

	fmt.Println()

	fmt.Println("Address of the 'userName' variable:", &userName)
	fmt.Println("Address of the 'userPhone' variable:", &userPhone)
}
