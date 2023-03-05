package utility

func ForLoop(number int) (sum int) {
	sum = 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum
}

func PressingStyle(a int, b float64) (result int) {
	result = a + int(b)
	return
}

func TwoSum(numbers []int, target int) [2]int {
	//   begin := numbers[0]
	for idx1, value := range numbers {
		for idx2, check := range numbers {
			if target-value == check && idx1 != idx2 {
				return [2]int{idx1, idx2}
			}
		}
	}
	return [2]int{0, 0}
}
