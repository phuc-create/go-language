package golesson

import (
	"fmt"
	"strconv"
	"time"
)

func Concurrency(animal string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 400)
	}
}
func middleWareChannel(c chan string, arr []string) {
	for i := 0; i < len(arr); i++ {
		strToNum, err := strconv.Atoi(arr[i])
		if err != nil {
			c <- arr[i]
		} else {
			numCombine := ""
			for i := 0; i <= strToNum; i++ {
				intToStr := strconv.Itoa(i)
				numCombine += intToStr
			}
			c <- numCombine
		}
	}
	close(c)
}

func DeadLockChecker() {

	// chanStop := make(chan int)
	arrayValues := []string{"Nguyen Huu Phuc", "5", "Phuc", "24", "22"}
	chanDeadLookCreation := make(chan string, len(arrayValues))

	go middleWareChannel(chanDeadLookCreation, arrayValues)
	for i := range chanDeadLookCreation {
		fmt.Println(i)
	}
}
