package main

import (
	"fmt"
	golesson "main/go-lesson"
	"main/utility"
)

// "os"

func main() {

	// Pressing style in Go lang
	fmt.Println("sum from 1 to 12: ", utility.ForLoop(12))
	golesson.DeadLockChecker()
}
