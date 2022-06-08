package main

import (
	golesson "main/go-lesson"
	"sync"
)

// "os"

func main() {
	// Checking function
	// fmt.Println(algorithms.TestBool())

	// Pressing style in Go lang
	// fmt.Println("Sum by pressing style: ", utility.PressingStyle(12, 32.902))
	// go routine
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		golesson.Concurrency("Dog")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
