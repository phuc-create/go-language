package algorithms

import (
	"regexp"
	"strings"
)

// My solution
func isVowel(c string) bool {
	vowels := strings.Split("ueoai", "")
	for _, v := range vowels {
		if c == v {
			return true
		}
	}
	return false
}

func Disemvowel(comment string) string {
	newString := ""
	slices := strings.Split(comment, "")

	for _, c := range slices {
		if !isVowel(strings.ToLower(c)) {
			newString += c
		}
	}
	return newString
}

// Other solution

func Disemvowel_Soluton_2(comment string) string {
	for _, c := range "aiueoAIUEO" {
		comment = strings.ReplaceAll(comment, string(c), "")
	}
	return comment
}

func Disemvowel_Soluton_3(comment string) string {
	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}

func Disemvowel_Soluton_4(comment string) string {
	vovels := map[string]bool{"a": true, "o": true, "e": true, "u": true, "i": true}
	str := ""

	for _, r := range comment {
		if !vovels[strings.ToLower(string(r))] {
			str = str + string(r)
		}
	}

	return str
}
