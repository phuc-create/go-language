package main

import (
	"fmt"
	"math"
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
	// Second way
	var variable2 string = "Variable Type String 2"
	fmt.Println(variable2)
	// Third way
	// You can call it "Type Inference"
	var variable3 = "Variable Type String 3"
	fmt.Println(variable3)
	// Last way
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
	// Another way
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
	// Short hand declaration
	language := "Go Lang"
	language2 := "Javascript"
	beginner := true
	fmt.Println("I am ", beginner, " at ", language, language2)
	// Checking function
	fmt.Println(testBool())
	// Types in Go
	var boolType bool = (true || false) && false // Boolean logic
	var stringType string = "hello world"        // String
	var intType int = 21                         // Integer - int8 int16 int32 int64
	var uintType uint = 1                        // Positive integer - uint8 uint16 uint32 uint64
	var byteType byte = 244                      // Byte - represent for uint8 0 -> 255
	var runeType rune = -2147483648              // Rune - represent for int32 -2147483648 –> 2147483647
	var complexType complex64 = 23 + 3i          // Complex Type
	com1 := complex(4, 7)
	com2 := 23 + 27i
	comAdd := com1 + com2
	fmt.Println("Add two complex we have ", comAdd)
	fmt.Println(boolType)
	fmt.Println(stringType)
	fmt.Println(intType)
	fmt.Println(uintType)
	fmt.Println(byteType)
	fmt.Println(runeType)
	fmt.Println(complexType)
	num1 := 20
	num2 := 30
	fmt.Println("Sum of 2 numbers ", num1, ",", num2, " la ", num1+num2)
	// Pressing style in Go lang
	fmt.Println("Sum by pressing style: ", pressingStyle(12, 32.902))
	// Min and Max of every type
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.Floor(2.7))    // Floor returns the greatest integer value less than or equal to x.
	fmt.Println(math.Ceil(2.7))     // Ceil returns the least integer value greater than or equal to x.
	fmt.Println(math.Max(342, 452)) // Find tha lagest value or number of tw numbers
	fmt.Println(math.Abs(-1319))    // negative number to positive
	// Working with byte
	var myByte byte = 255
	fmt.Printf("%T", myByte)
	fmt.Println()
	var asciiChrter byte = 'A'
	fmt.Println(asciiChrter)
}

func variables() string {
	var firstVariable = "Nguyen Huu Phuc"
	return "My name is " + firstVariable
}

func forLoop(number int) (sum int) {
	sum = 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return
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

func pressingStyle(a int, b float64) (result int) {
	result = a + int(b)
	return
}
