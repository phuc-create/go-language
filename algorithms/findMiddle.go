package algorithms

import "sort"

// My solution
func Gimme(array [3]int) int {
	n0 := array[0]
	n1 := array[1]
	n2 := array[2]
	if n1 < n0 && n0 < n2 || n1 > n0 && n0 > n2 {
		return 0
	}

	if n0 < n1 && n1 < n2 || n0 > n1 && n1 > n2 {
		return 1
	}
	return 2
}

// Other solution

func Gimme_Solution_2(array [3]int) int {
	sorted := make([]int, len(array))
	copy(sorted, array[:])
	sort.Ints(sorted)
	for k, v := range array {
		if v == sorted[1] {
			return k
		}
	}
	return -1
}

func Gimme_Solution_3(array [3]int) int {
	x := array
	sort.Ints(x[0:])
	for i, v := range array {
		if v == x[1] {
			return i
		}
	}
	return -1
}

func Gimme_Solution_4(array [3]int) int {
	x := make([]int, 3)

	copy(x, array[:])
	sort.Ints(x)

	for i, a := range array {
		if a == x[1] {
			return i
		}
	}

	return 0
}
