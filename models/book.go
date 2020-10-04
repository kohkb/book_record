package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title    string
	Category uint64
	Status   uint64
	Comment  string
}
