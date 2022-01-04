package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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
	// fmt.Println(math.MaxInt16)
	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.MaxInt64)
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
	const initValue = "Have a nice day!" // This is another way to declare variable by using const key word
	fmt.Println(initValue)

	// 3. If else condition
	num := 21
	if num == 21 {
		fmt.Println("Your age is 21")
	} else {
		fmt.Println("It's not your age!!!")
	}
	// Another way using if else statement
	if num := 21; num > 30 {
		fmt.Println("Your so old!")
	} else {
		fmt.Println("You so young!")
	}
	// Switch - case
	// Single value in one line
	yourname := "Phuc"
	switch yourname {
	case "Phuc":
		fmt.Println(yourname, "is your name")
	case "Sam":
		fmt.Println(yourname, "is also your name")
	default:
		fmt.Println("???")
	}
	hobbies := "Sleep"
	switch hobbies {
	case "Sleep", "Swim", "Soccer", "Basketball", "Tenis", "Mountain-Climb":
		fmt.Println("You are kind of amazing")
	default:
		fmt.Println("???")
	}
	// Fallthrough ,Break, Goto, Defer
	// 1. Fallthrough - Keep check all of condition in switch
	country := "VN"
	switch country {
	case "VN":
		fmt.Println("You are Viet Nam")
		fallthrough
	case "SP":
		fmt.Println("You are Singapore")
		fallthrough
	default:
		fmt.Println("???")
	}
	// 2. Break - check conditon in case and break if you don't want to keep tracking the next case
	number := 10
	switch number {
	case 10:
		if number == 10 {
			fmt.Println("Break here")
			break
		}
		fmt.Println("Number = 10")
		fallthrough
	case 20:
		fmt.Println("Number = 20")
		fallthrough
	default:
		fmt.Println("???")
	}
	// 2. Goto - if satisfy condition,let go to another check inside of case
	checker := 21
	switch checker {
	case 10:
		if number == 10 {
			fmt.Println("Break here")
			break
		}
		fmt.Println("Number = 10")
		fallthrough
	case 21:
		fmt.Println("Waiting...")
		if checker == 21 {
			goto ageAccepted // go to case ageAccepted

		}
	ageAccepted: // execute that
		fmt.Println(checker, "That is your age,congrat!!")
	default:
		fmt.Println("???")
	}
	// 4. Defer - A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("last output")
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
	array := []int{3, 100, 38, 10, 24, 1, 34}
	sort.Ints(array)
	fmt.Println(array)
	infors := somthing("Phuc", 21, "VN", true)
	fmt.Println(infors)
	TwoToOne("hello ửold", "ádasd")
}

func helpers() {
	panic("unimplemented")
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
func somthing(name string, age int, country string, gender bool) string {
	gen := func(g bool) string {
		if g {
			return "Man"
		}
		return "Girl"
	}(gender)
	lastAge := strconv.Itoa(age)
	return "Hey " + gen + ", your name is " + name + ",and you are " + lastAge + " years old, you from " + country
}

func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	fmt.Println(chars, chars[2])
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
	/* Another way
	func TwoToOne(s1 string, s2 string) string {
		result:=""
		s := strings.Split(s1+s2,"")
		sort.Strings(s)
		for _,c := range s {
			if result == "" || c != string(result[len(result)-1]) {
				result+=c
			}
		}
		return result
	}
	*/
}
