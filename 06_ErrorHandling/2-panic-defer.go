package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println("This is deferred1")
	defer fmt.Println("This is deferred2")
	panic("something bad happened")

	defer fmt.Println("This is deferred3")
	fmt.Println("end")

}

/*
NOTE:
	1. panic() is the last executing statement
	2. After panicking, it will execute all registered defers, till that line
*/
