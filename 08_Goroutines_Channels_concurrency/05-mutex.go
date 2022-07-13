package main

import (
	"fmt"
	"sync"
)

// mutex - needed for solving race condition
func main() {
	m := make(map[int]string)
	m[2] = "First Value"
	fmt.Println("Start - m =", m) //  map[2:First Value]

	var lock sync.Mutex
	go func() {
		lock.Lock()
		m[2] = "Second Value"
		lock.Unlock()
		fmt.Println("In - m =", m) //map[2:Second Value]
	}()

	lock.Lock()
	m[2] = "Third Value"
	lock.Unlock()
	fmt.Println("last - m =", m) //  map[2:Third Value]

	for {

	}

}
