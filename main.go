package main

import "fmt"

func employee(employeeID int, employeeName string) {
	fmt.Println("Employee ID:", employeeID)
	fmt.Println("Employee Name:", employeeName)
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a sample Go program.")
	employee(1, "John Doe")
	employee(2, "Jane Smith")
}
