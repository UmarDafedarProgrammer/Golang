package main

import "fmt"

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func getAllOddNumbers(nums []int) []int {
	var resultNums []int
	for _, num := range nums {
		if isOdd(num) == true { // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

func getAllEvenNumbers(nums []int) []int {
	var resultNums []int
	for _, num := range nums {
		if isEven(num) == true { // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

// Decorator design pattern - passing func as argument
type EvenOddFunc func(int) bool

func Filter(nums []int, myfunc EvenOddFunc) []int {
	var resultNums []int
	for _, num := range nums {
		if myfunc(num) == true { // function
			resultNums = append(resultNums, num)
		}
	}
	return resultNums
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7, 8, 9}
	fmt.Println("slice = ", slice)

	fmt.Println("isOdd(2) =", isOdd(2))
	fmt.Println("isOdd(3) =", isOdd(3))
	fmt.Println()

	for _, num := range slice {
		fmt.Printf("isOdd(%d) =%t\n", num, isOdd(2))
	}

	for _, num := range slice {
		fmt.Printf("isEven(%d) =%t\n", num, isEven(2))
	}

	fmt.Println("getAllOddNumbers(slice)  =", getAllOddNumbers(slice))
	fmt.Println("getAllEvenNumbers(slice) =", getAllEvenNumbers(slice))
	fmt.Println()

	fmt.Println("Filter(slice, isOdd)     =", Filter(slice, isOdd))
	fmt.Println("Filter(slice, isEven)    =", Filter(slice, isEven))
	fmt.Println("Filter(slice, isEven)    =", Filter(slice, isPostive))
}

func isPostive(num int) bool {
	if num >= 0 {
		return true
	}
	return false
}
