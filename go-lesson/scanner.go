package golesson

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type somthing: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Now you are %d years old\n", 2021-input)
}
