package main

import (
	"fmt"
	"sync"
	"time"
)

/*
To wait for multiple goroutines to finish, we can use a wait group.

sync.WaitGroup is struct type with three methods:
	- Add: It adds an integer counter number, that says main function
		   will wait for that number of goroutines to complete.
	- Done: It decreases the count, which is added in Add() method
		- It should be called at the end of goroutine
	- Wait: It waits for all goroutines to finish.
		- When count reaches 0, it finishes the wait.
*/

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worder:%d - Started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("\tWorder:%d - Finished\n", id)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(0, &wg)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

/*
~go run 02_waitgroups.go
Worder:5 - Started
Worder:2 - Started
Worder:0 - Started
Worder:3 - Started
Worder:4 - Started
Worder:1 - Started
        Worder:1 - Finished
        Worder:3 - Finished
        Worder:0 - Finished
        Worder:5 - Finished
        Worder:4 - Finished
        Worder:2 - Finished



*/
