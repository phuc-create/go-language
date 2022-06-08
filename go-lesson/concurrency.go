package golesson

import (
	"fmt"
	"time"
)

func Concurrency(animal string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 400)
	}
}
