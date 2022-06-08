package golesson

import "fmt"

func TypeOfGoLang() {
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
}
