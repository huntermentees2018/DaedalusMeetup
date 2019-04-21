package models

import "github.com/jinzhu/gorm"

type Met struct {
	gorm.Model
	studentOneID int
	studentTwoID int
}
