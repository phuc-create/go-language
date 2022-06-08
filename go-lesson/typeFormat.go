package golesson

import "fmt"

func TypeFormat() {
	// %T => Type
	// %s => String
	// %q => Double quote string
	// %v => Value of variable default format
	// %t => True or False
	// %b => base2 (for interger only)
	// %o => base8
	// %d => base10
	// %x => base16
	// %f => default width & default precision
	// %9f => 9 width & default precision
	// %.2f => default width & 2 precision
	// %9.2f => 9 width & 2 precision
	// %09d => padding digits t length 9 with preceeding 0's
	// %-9d => padding with spaces (width 4,left justified)
	fmt.Printf("Number: %-5q  hello world", "Phuc")
}
