package golesson

import "fmt"

func example(name string) {
	fmt.Println(name)
}
func GoRoutine() {
	// Go routine represent for a waiting, if the function call with go prefix => will execute stack by stack and because the function main required a cuple of time to wait till the go routine execute successfullly so we need to use some option bellow to make sure the result can be returned
	// #1 - go prefix
	go example("Sam")
	// #2 -  wait group

}
