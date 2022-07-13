package main

import (
	"fmt"
)

func main() {

	mystr := "Srinivas"

	/*
		S r i n i v a s
		0 1 2 3 4 5 6 7

		string len - 8
		8 -3 = 5
	*/
	// SYNTAX  : str[starIndex: finalindex]
	/*
		default
			start index = 0
			final index = len(string)
	*/
	fmt.Println(mystr)                      // Srinivas
	fmt.Println(mystr[1], string(mystr[1])) // 114 r
	fmt.Println(mystr[5], string(mystr[5])) // 118 v

	fmt.Println(mystr[2:5]) // ini
	fmt.Println(mystr[5:7]) // va
	fmt.Println(mystr[5:8]) // vas

	// fmt.Println(mystr[5:9])
	// panic: runtime error: slice bounds out of range [:9] with length 8

	// how to get 'vas'  - last three chars

	fmt.Println(len(mystr)) // 8
	fmt.Println(mystr[5:8], mystr[5:])
	fmt.Println(len(mystr) - 3) // 5

	fmt.Println(mystr[5], mystr[len(mystr)-3])

	fmt.Println(mystr[5:8], mystr[len(mystr)-3:8])
	fmt.Println(mystr[5:], mystr[len(mystr)-3:])

	fmt.Println()
	fmt.Println(mystr[:5]) // mystr[0:5]
	fmt.Println(mystr[2:]) // mystr[2:8]
	fmt.Println(mystr[:])  // mystr[0:8]

}
