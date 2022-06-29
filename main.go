package main

import (
	"fmt"
	golesson "main/go-lesson"
	"main/utility"
	"net/http"
	// "strconv"
	// "sync"
)

// "os"

func main() {

	// Pressing style in Go lang
	fmt.Println("sum from 1 to 12: ", utility.ForLoop(12))
	// golesson.DeadLockChecker()
	// fmt.Println(rand.Intn(10)) RANDOM NUMBER
	// a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	// a := []int{-10000000, 100000000}
	// b := []int{100000000000000, 10000000000000000}
	// fmt.Println(algorithms.Comp2(a, b))
	person := golesson.StrucDeclare()
	fmt.Println(person)
	http.HandleFunc("/hello", golesson.HelloConnecter)
	http.HandleFunc("/headers", golesson.HelloHeader)

	// func Closure() func() int {
	// 	i := 0
	// 	return func() int {
	// 		i++
	// 		return i
	// 	}
	// }
	// closeExample := golesson.Closure()
	golesson.RunTest()
	//
	// 	fmt.Println(closeExample())
	// 	fmt.Println(closeExample())
	// 	fmt.Println(closeExample())

	// http.ListenAndServe(":7070", nil)
}
