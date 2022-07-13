package main

import (
	"fmt"
	"unicode/utf8"
)

// go install golang.org/x/text/cases@latest
func main() {
	s := "Hello, 世界"
	fmt.Println("string length:", len(s))                    // "13"
	fmt.Println("rune count   :", utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); i++ {
		fmt.Print("\n", string(s[i]), "   ", s[i])
		r, size := utf8.DecodeRuneInString(string(s[i]))
		fmt.Println("   r=", r, "size=", size)
	}

	str := []byte("Hello12! दुनिया ಜಗತ್ತು ప్రపంచ 世界") 

	for len(str) > 0 {
		r, size := utf8.DecodeRune(str)
		fmt.Printf("%c %v bytes\n", r, size)

		str = str[size:]
	}
	fmt.Println()

	valid := []byte("Hello, ప్రపంచ") // World in Telugu
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))   // true
	fmt.Println(utf8.Valid(invalid)) // false

	// c := cases.Fold()
	// fmt.Printf("%s %v", c, c.String("grüßen"))

}

// ß   ss
