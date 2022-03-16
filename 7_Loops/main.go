package main

import (
	"fmt"
	"strings"
)

func main() {
	// for i := 0; i < 4; i++ {
	// 	fmt.Println("Hello World")
	// }

	var name = [5]string{"Tanay Pratap", "Rahul Pandey", "Arnav Gupta", "Kunal Kushwaha", "Eddie Jaoude"}

	firstNames := []string{}

	for _, fullName := range name {
		var names = strings.Fields(fullName)
		firstNames = append(firstNames, names[0])
	}

	fmt.Print(firstNames)
}
