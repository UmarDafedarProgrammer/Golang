package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

func partialPlus(x int) func(int) int {
	return func(y int) int {
		return plus(x, y)
	}
}

func addition(n1 float32) func(float32) float32 {
	return func(n2 float32) float32 {
		fmt.Println("\tn1 = ", n1)
		fmt.Println("\tn2 = ", n2)
		return n1 + n2
	}
}

func main() {

	// myfunc(1, 4)
	// myfunc(1, 3)
	// myfunc(1, 5)
	// myfunc(1, 43)
	// myfunc(1, 324)
	// myfunc(1, 2134)

	plusOne := partialPlus(1)
	fmt.Println("plusOne(5):", plusOne(5)) //prints 6
	fmt.Println("plusOne(7):", plusOne(7)) //prints 8
	fmt.Println("plusOne(8):", plusOne(8)) //prints 9

	fmt.Println()
	addWith10 := partialPlus(10)
	fmt.Println("addWith10(5):", addWith10(5)) //prints 15
	fmt.Println("addWith10(7):", addWith10(7)) //prints 17
	fmt.Println("addWith10(8):", addWith10(8)) //prints 18

	fmt.Println()
	plusFive := addition(5)
	fmt.Println("plusFive(10):", plusFive(10))   // 15
	fmt.Println("plusFive(12):", plusFive(12))   // 15
	fmt.Println("plusFive(-10):", plusFive(-10)) // -5
	fmt.Println("plusFive(2.3):", plusFive(2.3)) // 7.3
}
