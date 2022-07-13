package main

import "fmt"

func main() {
	var number1 int
	fmt.Println("Before, number1=", number1) // 0

	fmt.Print("Enter a number:")
	fmt.Scanf("%d", &number1)
	fmt.Println("After, number1=", number1)

}

/*

~go run p-runtime-input.go
Before, number1= 0
Enter a number:123
After, number1= 123

~go run p-runtime-input.go
Before, number1= 0
Enter a number:543.65465
After, number1= 543

~
~go run p-runtime-input.go
Before, number1= 0
Enter a number:4323.65465.765
After, number1= 4323

~go run p-runtime-input.go
Before, number1= 0
Enter a number:my123
After, number1= 0

~go run p-runtime-input.go
Before, number1= 0
Enter a number:54365hfhfh
After, number1= 54365

~go run p-runtime-input.go
Before, number1= 0
Enter a number:true
After, number1= 0

~go run p-runtime-input.go
Before, number1= 0
Enter a number:nil
After, number1= 0

~



*/
