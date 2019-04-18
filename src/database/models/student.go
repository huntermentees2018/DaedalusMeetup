package models

import "github.com/jinzhu/gorm"

// Student struct
type Student struct {
	gorm.Model
	Name  string
	Email string
}
