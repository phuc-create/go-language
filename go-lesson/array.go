package golesson

import (
	"fmt"
	"sort"
)

func ArrayInGo() {
	array := []int{3, 100, 38, 10, 24, 1, 34}
	sort.Ints(array)
	fmt.Println(array)
	// below is practice :)))))
	infors := Somthing("Phuc", 21, "VN", true)
	fmt.Println(infors)
}
