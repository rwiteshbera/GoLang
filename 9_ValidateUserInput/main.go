package main

import (
	"fmt"
	"strings"
)

func main() {
	// Validate user input
	var firstName string
	var lastName string
	var email string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")

	if isValidName && isValidEmail {
		fmt.Print("\nSuccessfully Registered\n")
	} else {
		fmt.Printf("\nSomething went wrong. Make sure you have provided correct details.\n")
	}

}
