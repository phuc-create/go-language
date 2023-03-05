package golesson

import "fmt"

func Slices() {
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

	// The type []T is a slice with elements of type T.
	primes := []int{2, 3, 5, 7, 11, 13, 23, 31, 123, 12, 123, 23}

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:

	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.

	var s []int = primes[1:7] // a[1:7]
	// The following expression creates a slice which includes elements 1 through 7 of a:
	fmt.Println(s) // [3 5 7 11 13 23]

}
