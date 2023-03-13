package golesson

// Or you can declare an interface there

type Number interface {
	int | int32 | int64 | float32 | float64
}

func SimpleCalculatorFuncUseGenericsType[T Number](input []T) T {
	var sum T = 0
	for _, num := range input {
		sum += num
	}
	return sum
}
