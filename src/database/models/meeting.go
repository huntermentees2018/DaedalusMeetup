package models

import "github.com/jinzhu/gorm"

// Meeting describes a meeting between two students with the corresponding event id and date that it is happening
type Meeting struct {
	gorm.Model
	StudentOneID int
	StudentTwoID int
	EventID      string
	DateTime     string
}
