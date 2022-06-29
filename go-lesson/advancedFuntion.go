package golesson

import "fmt"

func test(Func func(int) int) { // a function have params is function and return function call
	sum := Func(6)   //call param as function
	fmt.Println(sum) // print the value after function called
}

func closure(x string) func() { // closure return a function which will take param of function itself
	return func() { // handle function inside closure and using params of parent function if needed
		fmt.Println(x) // print out the value of closure function
	}
}

func RunTest() {
	sum := func(x int) int {
		return x * x
	}

	cls := closure

	cls("Hello")()

	test(sum)
}
