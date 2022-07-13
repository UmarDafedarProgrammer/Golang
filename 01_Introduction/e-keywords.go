package main

import "fmt"

func main() {
	fruit := "apple"
	fmt.Println("fruit")
	fmt.Println(fruit)

	// NOTE 1: keywords should not be used as identifiers
	// break := "one" //syntax error: unexpected := at end of statement

	breAK := "two"
	fmt.Println("breAK =", breAK)

	// NOT RECOMMENDED - predeclared variables should not be used as keywords
	true := "One"
	fmt.Println("true:", true)

	true1 := "One"
	fmt.Println("true1:", true1)

	break_one := "one"
	fmt.Println("break_one =", break_one)

	// NOTE 2: CamelCase is recommended for identifiers

	breakOne := "one"
	fmt.Println("breakOne=", breakOne)

	costOfMangos := 34
	fmt.Println("costOfMangos=", costOfMangos)

	NoOfProcessesRunning := 3
	fmt.Println("NoOfProcessesRunning=", NoOfProcessesRunning)

	NoOfValidVoters := 600
	fmt.Println("NoOfValidVoters ->", NoOfValidVoters)

}

// NOTE:  go is a case-sensitive language

// :=  walrus operator

// variable
// identifiers or user-defined variable
// keywords  -- language builtin words
