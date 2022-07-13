package main

// Methods with Pointer receivers as Functions

import (
	"fmt"
)

type PointA struct {
	X, Y float64
}

// Pointer receivers as function argument
func AddDistance(p1 PointA, dx, dy float64) {
	// call by value
	p1.X = p1.X + dx
	p1.Y += dy

	fmt.Printf("\nIn AddDistance, %+v\n", p1) // {X:13 Y:9}
}

func AddDistancePtr(p2 *PointA, dx, dy float64) {
	// call by reference
	p2.X = p2.X + dx
	p2.Y += dy
	fmt.Printf("\nIn AddDistancePtr, %+v\n", p2) // &{X:13 Y:9}

}


func main() {
	p := PointA{3, 4}
	fmt.Printf("Point p = %+v\n", p) //   {X:3 Y:4}

	// Using functions
	AddDistance(p, 10, 5)
	fmt.Printf("After AddDistance, %+v\n", p) // {X:3 Y:4}

	AddDistancePtr(&p, 10, 5)
	fmt.Printf("After AddDistancePtr, %+v\n", p) // {X:13 Y:9}


}
