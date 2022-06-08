package golesson

import (
	"errors"
	"fmt"
	"main/utility"
	"reflect"
)

func Calculate() {
	var test1 float64 = 457
	var test2 int = 3
	fmt.Println("Devide is:", test1/float64(test2))
	fmt.Println("Result of for loop through 10 is", utility.ForLoop(10))
}

func SimpleCal(a int, b float64) (int, error) {
	typeOfa := reflect.TypeOf(a).Kind()
	typeOfb := reflect.TypeOf(b).Kind()

	if typeOfa == reflect.Int && typeOfb == reflect.Float64 || typeOfb == reflect.Int {
		sum := a + int(b)
		return sum, nil
	}
	return 0, errors.New("use correct type for number")
}
