package algorithms

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	fmt.Println(chars, chars[2])
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
	/* Another way
	func TwoToOne(s1 string, s2 string) string {
		result:=""
		s := strings.Split(s1+s2,"")
		sort.Strings(s)
		for _,c := range s {
			if result == "" || c != string(result[len(result)-1]) {
				result+=c
			}
		}
		return result
	}
	*/
}
