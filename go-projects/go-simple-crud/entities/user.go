package entities

import "fmt"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id:%s,\nusername:%s\npassword:%s", user.Id, user.Username, user.Password)
}
