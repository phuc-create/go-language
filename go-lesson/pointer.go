package golesson

import (
	"fmt"
)

type Information struct {
	Email    string
	Password *string
}

var passConfig string = "Phuc1492"

var infor = Information{
	Email:    "huuphuc0006@gmail.com",
	Password: &passConfig,
}

func (inf *Information) UpdateEmail(email string) {
	inf.Email = email
}

func (inf *Information) UpdatePassword(password string) {
	inf.Password = &password
}

func PointerLesson(value int) {
	infor.UpdateEmail("new@gmail.com")
	infor.UpdatePassword("newpass")
	fmt.Printf("Your email is: %v\nYour password is: %v\n", infor.Email, *infor.Password)
}

/*
EXAMPLE ABOUT POINTER
package main

import "fmt"

func main() {
	var a = 5
	var p = &a

	fmt.Println("Giá trị được lưu trữ trong biến a = ", a) // => 5
	fmt.Println("Địa chỉ của biến a = ", &a) // => 0xc00001e0d0
	fmt.Println("Giá trị được lưu trữ trong biến p = ", p) // => 0xc00001e0d0
}
*/
