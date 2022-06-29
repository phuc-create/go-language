package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

type Book struct {
	gorm.Model
	ID     string `json:"name"`
	Author string `json:"author"`
	Public bool   `json:"public"`
}
