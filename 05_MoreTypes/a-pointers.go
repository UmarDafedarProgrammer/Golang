package main

import "fmt"

/*
pointers are references to defined variables
'call by reference' can be done based on pointer

*/

func main() {

	var num1 int = 13 /* actual variable declaration */

	fmt.Printf("\nvalue            : %v", num1)  // 13
	fmt.Printf("\ntype             : %T", num1)  // int
	fmt.Printf("\naddress location : %x", &num1) // c0000100a8

	// every variable is a memory location and every memory location has its address
	// pointer is a variable whose value is the address of another variable,
	// i.e., direct address of the memory location.
	fmt.Println()

	var num1_ptr *int /* pointer variable declaration */

	fmt.Printf("\nvalue            : %v", num1_ptr) // <nil>
	fmt.Printf("\ntype             : %T", num1_ptr) // *int
	fmt.Printf("\naddress location : %x", num1_ptr) // 0
	fmt.Println()

	/* store address of a in pointer variable*/
	num1_ptr = &num1
	fmt.Printf("\nvalue            : %v", num1_ptr) // 0xc0000100a8
	fmt.Printf("\ntype             : %T", num1_ptr) // *int
	fmt.Printf("\naddress location : %x", num1_ptr) // c0000100a8
	fmt.Println()

	// How to access the value referenced by the pointer
	fmt.Println("\n*num1_ptr    =", *num1_ptr)

	fmt.Println("num1 == *num1_ptr :", num1 == *num1_ptr) // value
	fmt.Println("&num1 == num1_ptr :", &num1 == num1_ptr) // address
}

/*  & - ampercent

num1
	- value 13
	- address : 123

num1_ptr
	- relation
			num1_ptr = &num1






*/
