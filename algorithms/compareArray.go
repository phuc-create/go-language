// Given two arrays a and b write a function comp(a, b) (orcompSame(a, b)) that checks whether the two arrays have the "same" elements, with the same multiplicities (the multiplicity of a member is the number of times it appears). "Same" means, here, that the elements in b are the elements in a squared, regardless of the order.

// Examples
// Valid arrays
// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
// comp(a, b) returns true because in b 121 is the square of 11, 14641 is the square of 121, 20736 the square of 144, 361 the square of 19, 25921 the square of 161, and so on. It gets obvious if we write b's elements in terms of squares:

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
// Invalid arrays
// If, for example, we change the first number to something else, comp is not returning true anymore:

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
// comp(a,b) returns false because in b 132 is not the square of any number of a.

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
// comp(a,b) returns false because in b 36100 is not the square of any number of a.

// Remarks
// a or b might be [] or {} (all languages except R, Shell).
// a or b might be nil or null or None or nothing (except in C++, COBOL, Crystal, D, Dart, Elixir, Fortran, F#, Haskell, Nim, OCaml, Pascal, Perl, PowerShell, Prolog, PureScript, R, Racket, Rust, Shell, Swift).
// If a or b are nil (or null or None, depending on the language), the problem doesn't make sense so return false.

package algorithms

import (
	"reflect"
	"sort"
)

// Nhan sensei solution
func Comp(array1 []int, array2 []int) bool {
	if (array1 == nil) || (array2 == nil) {
		return false
	}
	if len(array1) != len(array2) {
		return false
	}

	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for index, _ := range array1 {
		key1 := array1[index] * array1[index]
		key2 := array2[index]

		map1[key1] = map1[key1] + 1
		map2[key2] = map2[key2] + 1

	}

	for key, value := range map1 {
		if value2, err := map2[key]; !err {
			return false
		} else if value2 != value {
			return false
		}
	}
	return true
}

//Sam solution
func Comp2(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || (len(array1) != len(array2)) {
		return false
	}

	for i, value := range array1 {
		num := value * value
		array1[i] = num
	}

	sort.Ints(array2[:])
	sort.Ints(array1[:])
	return reflect.DeepEqual(array1, array2)
}

// It("should handle basic cases", func() {
// 	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
// 	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
// 	Comp(a1, a2, true) expected true
// 	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
// 	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
// 	Comp(a1, a2, false) expected false
// 	a1 = nil
// 	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
// 	Comp(a1, a2, false) expected false
// })
