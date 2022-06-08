package golesson

import "fmt"

func ByteTypeInGo() {
	// Working with byte
	var myByte byte = 255
	fmt.Printf("%T", myByte)
	fmt.Println()
	var asciiChrter byte = 'A'
	fmt.Println(asciiChrter)
	const initValue = "Have a nice day!" // This is another way to declare variable by using const key word
	fmt.Println(initValue)
}
