package main

// Methods with Pointer receivers as Functions

import (
	"fmt"
)

type PointA struct {
	X, Y float64
}


// Method
func (p3 PointA) AddDistanceMd(dx, dy float64) {
	// call by value
	p3.X = p3.X + dx
	p3.Y += dy
	fmt.Printf("\nIn AddDistanceMd, %+v\n", p3) // {X:20 Y:10}
}

func (p4 *PointA) AddDistanceMdPtr(dx, dy float64) {
	// call by reference
	p4.X = p4.X + dx
	p4.Y += dy
	fmt.Printf("\nIn AddDistanceMdPtr, %+v\n", p4) // &{X:20 Y:10}
}

func main() {
	p := PointA{3, 4}
	fmt.Printf("Point p = %+v\n", p) //   {X:3 Y:4}

	// Using method
	p.AddDistanceMd(7, 1)
	fmt.Printf("After p.AddDistanceMd, %+v\n", p) // {X:13 Y:9}

	p.AddDistanceMdPtr(7, 1)
	fmt.Printf("After p.AddDistanceMdPtr, %+v\n", p) // {X:20 Y:10}

}
