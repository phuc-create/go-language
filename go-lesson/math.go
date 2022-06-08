package golesson

import (
	"fmt"
	"math"
)

func MathInGo() {

	// Min and Max of every type
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.Floor(2.7))    // Floor returns the greatest integer value less than or equal to x.
	fmt.Println(math.Ceil(2.7))     // Ceil returns the least integer value greater than or equal to x.
	fmt.Println(math.Max(342, 452)) // Find tha lagest value or number of tw numbers
	fmt.Println(math.Abs(-1319))    // negative number to positive
}
