package utility

func ForLoop(number int) (sum int) {
	sum = 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}

func PressingStyle(a int, b float64) (result int) {
	result = a + int(b)
	return
}
