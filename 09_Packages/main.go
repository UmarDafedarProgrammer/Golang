package main

import (
	b "employee/basic"
	"employee/basic/gross"

	"fmt"
)

func main() {
	// b.Basic = 1000

	fmt.Print("Enter your Basic Salary (in rupees):")
	fmt.Scanf("%d", &b.Basic)

	fmt.Println("gross.GrossSalary() = ", gross.GrossSalary())

}
