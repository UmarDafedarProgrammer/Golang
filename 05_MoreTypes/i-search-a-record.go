package main

import (
	"fmt"
	"strings"
)

type citySize struct {
	Name       string
	Population float64
}

func main() {
	// cities with the largest population in the world
	var citiesList = []citySize{
		{"Tokyo", 38001000},
		{"Delhi", 25703168},
		{"Shanghai", 23740778},
		{"Sao Paulo", 21066245},
	}

	searchCity := "shanghai"
	for _, city := range citiesList {
		// convert to lower case for exact matching
		if strings.ToLower(city.Name) == strings.ToLower(searchCity) {
			fmt.Println("WE found")
		}
	}
}

// Assignment
// 1. get city name in runtime, and get this poplation if present
// 2. Get population count, and return cities whose population is greater than given count
