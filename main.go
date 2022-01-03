package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(variables())
	fmt.Println("Result of for loop through 10 is", forLoop(10))
	// Declare Variables in Go lang
	// We have 4 ways to declare any variable in go by following these ways
	// first way
	var variable string
	variable = "Variable Type String"
	fmt.Println(variable)
	//second way
	var variable2 string = "Variable Type String 2"
	fmt.Println(variable2)
	//third way
	//you can call it "Type Inference"
	var variable3 = "Variable Type String 3"
	fmt.Println(variable3)
	//last way
	variable4 := "Variable Type String 4"
	fmt.Println(variable4)
	// Declare more than 1 variable in go
	// 1. Declare variables with same type

	var a, b int
	a = 10
	b = 20
	fmt.Println(a, b)
	var a2, b2 int = 10, 11
	fmt.Println(a2, b2)
	var a3, b3 int = 20, 21
	fmt.Println(a3, b3)

	// 2. Declare variables with different type
	var (
		name    string
		age     int
		address string
	)

	name = "Phuc"
	age = 21
	address = "53,Khue Ngoc Dien"
	fmt.Println(name, age, ",", address)
	//another way
	var (
		name2    string = "Phuc"
		age2     int    = 21
		address2 string = "53,Khue Ngoc Dien"
	)
	fmt.Println(name2, age2, ",", address2)

	var (
		name3    = "Phuc"
		age3     = 21
		address3 = "53,Khue Ngoc Dien"
	)
	fmt.Println(name3, age3, ",", address3)

	var name4, age4, address4 = "Sam", 21, "53,Khue Ngoc Dien"

	fmt.Println(name4, age4, ",", address4)
	//Short hand declaration
	language := "Go Lang"
	language2 := "Javascript"
	beginner := true
	fmt.Println("I am ", beginner, " at ", language, language2)
	//checking function
	fmt.Println(testBool())
}

func variables() string {
	var firstVariable = "Nguyen Huu Phuc"
	return "My name is " + firstVariable
}

func forLoop(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}

func testBool() int {
	variable := -5
	result := func(checker bool, acceptVal int, ignoreVal int) int {
		if checker {
			return acceptVal
		}
		return ignoreVal
	}(variable > 0, variable, -variable)
	return result
}
