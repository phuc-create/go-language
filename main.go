package main

import (
	golesson "main/go-lesson"
)

// "strconv"
// "sync"

// "os"

func main() {
	// golesson.GetPersonInterface()
	// 	per1 := golesson.BasicInformation{UUID: "111", Name: "sam", Age: 23, Phone: "0123123"}
	// 	per2 := golesson.BasicInformation{UUID: "222", Name: "torento", Age: 24, Phone: "234234"}
	// 	per3 := golesson.BasicInformation{UUID: "333", Name: "sam", Age: 21, Phone: "112909"}
	// 	extra1 := golesson.ExtraInformation{UUID: "111", Address: "234/123/1 Aurocoer"}
	// 	extra2 := golesson.ExtraInformation{UUID: "333", Address: "1 Minentan, District 20/2223"}
	//
	// 	pers := []golesson.PersonInterface{per1, per2, per3, extra1, extra2}
	// 	golesson.CombineBasedOnUuid(pers)

	// 	sq := golesson.Square{X: 5, Y: 6} //30
	// 	tg := golesson.Triangle{B: 5, C: 6}
	//
	// 	fmt.Println(sq.Area())
	// 	fmt.Println(tg.Area())

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

	// taskFn := func(ctx context.Context) <-chan int {
	// 	dst := make(chan int)
	// 	n := 1
	// 	go func() {
	// 		for {
	// 			select {
	// 			case <-ctx.Done():
	// 				fmt.Println("exe cancel")
	// 				return // returning not to leak the go routine
	// 			case dst <- n:
	// 				n++
	// 			}
	// 		}
	// 	}()
	// 	return dst
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // cancel when we are finished consuming integers
	// for n := range taskFn(ctx) {
	// 	fmt.Println(n)
	// 	if n == 5 {
	// 		break
	// 	}
	// }
	// var newString *string = new(string)   // returns a pointer to a newly allocated => 0xc00005c0d0
	// var useMap2 *[]string = new([]string) // returns a newly allocated array => &[]
	// newlyString := newString
	// *newString = "this is the value of newlyString also"
	// fmt.Println(*newString)
	// fmt.Println(*newlyString)
	// //------------------
	// fmt.Println(useMap2)
	// fmt.Println(*useMap2)

	// sum := golesson.SimpleCalculatorFuncUseGenericsType([]int{1, 2, 4, 3, 2})
	// fmt.Println(sum)
}

// func count(name string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- name + "-" + strconv.Itoa(i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }
