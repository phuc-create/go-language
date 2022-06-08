package golesson

import "fmt"

func LoopInGo() {
	// Loops
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	// Break and continue in loops
	// 1. Break
	for age := 0; age < 30; age++ {
		if age == 21 {
			fmt.Println("\n", age)
			break
		}
	}
	// 2. Continue
	for age := 0; age < 30; age++ {
		if age != 21 {
			continue
		} else {
			fmt.Println(age)
		}
	}
	for i, j := 0, 0; i <= 5 && j < 6; i, j = i+1, j+1 {
		if i == 5 {
			fmt.Println(i, j)
		}
	}
}
