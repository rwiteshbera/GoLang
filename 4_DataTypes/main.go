package main

import "fmt"

func main() {
	// Declaraing datatype of the variable
	var employee string
	var employeeAge int

	employee = "Ishan"
	employeeAge = 35
	fmt.Println("Employee Name:", employee)
	fmt.Printf("%v is %v years old.\n", employee, employeeAge)

	// %T prints type of the variable
	fmt.Printf("employee is %T and employeeAge is %T\n", employee, employeeAge)

	// Floating Point Number
	var salary float32 = 20000.00
	fmt.Println("Salary", salary)

	// Boolean Data Types
	var c bool = true
	fmt.Print(c)

}
