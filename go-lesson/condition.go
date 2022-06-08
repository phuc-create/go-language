package golesson

import (
	"fmt"
	"strconv"
)

func ConditionInGo() {
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
}

func SwitchCaseInGo() {
	// Switch - case
	// ------------------------------------------------------
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
	// ------------------------------------------------------
	// mutilple values in one line
	hobbies := "Sleep"
	switch hobbies {
	case "Sleep", "Swim", "Soccer", "Basketball", "Tenis", "Mountain-Climb":
		fmt.Println("You are kind of amazing")
	default:
		fmt.Println("???")
	}
	// ------------------------------------------------------
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
	// ------------------------------------------------------
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
	// ------------------------------------------------------
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
	// ------------------------------------------------------
	// 4. Defer - A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("last output")
}

func Somthing(name string, age int, country string, gender bool) string {
	gen := func(g bool) string {
		if g {
			return "Man"
		}
		return "Girl"
	}(gender)
	lastAge := strconv.Itoa(age)
	return "Hey " + gen + ", your name is " + name + ",and you are " + lastAge + " years old, you from " + country
}
