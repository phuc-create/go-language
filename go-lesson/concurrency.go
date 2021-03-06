package golesson

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Concurrency(animal string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 400)
	}
}

func middleWareChannel(c chan string, arr string) {
	defer wg.Done()
	strToNum, err := strconv.Atoi(arr)
	if err != nil {
		c <- arr
	} else {
		numCombine := ""
		for i := 0; i <= strToNum; i++ {
			intToStr := strconv.Itoa(i)
			numCombine += intToStr
		}
		c <- numCombine
	}
}

func DeadLockChecker() {

	// chanStop := make(chan int)
	arrayValues := []string{"Nguyen Huu Phuc", "5", "Phuc", "24", "22"}
	chanDeadLookCreation := make(chan string, len(arrayValues))

	for _, value := range arrayValues {
		wg.Add(1)
		go middleWareChannel(chanDeadLookCreation, value)
	}

	for i := 0; i < len(arrayValues); i++ {
		select {
		case value, ok := <-chanDeadLookCreation:
			if !ok {
				chanDeadLookCreation = nil
			}
			fmt.Println(value)
		}
		if chanDeadLookCreation == nil {
			break
		}
	}

	go func() {
		defer close(chanDeadLookCreation)
		wg.Wait()
	}()
}

func ConcurrencyAnotherExample() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
