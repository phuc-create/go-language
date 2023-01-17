package main

import (
	"context"
	"fmt"
	golesson "main/go-lesson"
)

// "strconv"
// "sync"

// "os"

func main() {
	sq := golesson.Square{X: 5, Y: 6} //30
	tg := golesson.Triangle{B: 5, C: 6}

	fmt.Println(sq.Area())
	fmt.Println(tg.Area())

	// Pressing style in Go lang
	// fmt.Println("sum from 1 to 12: ", utility.ForLoop(12))
	// golesson.DeadLockChecker()
	// fmt.Println(rand.Intn(10)) RANDOM NUMBER
	// a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	// a := []int{-10000000, 100000000}
	// b := []int{100000000000000, 10000000000000000}
	// fmt.Println(algorithms.Comp2(a, b))
	// person := golesson.StrucDeclare()
	// fmt.Println(person)
	// http.HandleFunc("/hello", golesson.HelloConnecter)
	// http.HandleFunc("/headers", golesson.HelloHeader)

	// func Closure() func() int {
	// 	i := 0
	// 	return func() int {
	// 		i++
	// 		return i
	// 	}
	// }
	// closeExample := golesson.Closure()
	// golesson.RunTest()
	//
	// 	fmt.Println(closeExample())
	// 	fmt.Println(closeExample())
	// 	fmt.Println(closeExample())

	// http.ListenAndServe(":7070", nil)
	// fmt.Println(TwoSum([]int{1, 2, 3}, 4))
	// fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690))
	// fmt.Println(TwoSum([]int{2, 2, 3}, 4))
	// c := make(chan string)
	// c2 := make(chan string)
	// go count("goroutine", c)
	// // count("concurency", c2)
	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	// time.Sleep(time.Second * 2000)
	x := "Hello world"
	var y *string = &x
	*y = "Hi Mom"
	// golesson.PointerLesson(3)
	golesson.Process()
	// fmt.Printf("Type of x is: %T\n", x)
	// changeValue(&x)
	// fmt.Printf("Value of y is: %v\n", y)
	// fmt.Printf("Reference value of y is: %v\n", *y)
	// fmt.Println("Address of x is:", &x)
	taskFn := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("exe cancel")
					return // returning not to leak the go routine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	for n := range taskFn(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}

// func count(name string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- name + "-" + strconv.Itoa(i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }

func TwoSum(numbers []int, target int) [2]int {
	//   begin := numbers[0]
	for idx1, value := range numbers {
		for idx2, check := range numbers {
			if target-value == check && idx1 != idx2 {
				return [2]int{idx1, idx2}
			}
		}
	}
	return [2]int{0, 0}
}
