package main

import (
	"fmt"
	"main/algorithms"
	"main/utility"
)

// "os"

func main() {

	// Pressing style in Go lang
	fmt.Println("sum from 1 to 12: ", utility.ForLoop(12))
	//golesson.DeadLockChecker()
	// a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	a := []int{-10000000, 100000000}
	b := []int{100000000000000, 10000000000000000}
	fmt.Println(algorithms.Comp2(a, b))
}
