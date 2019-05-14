package models

import (
	"github.com/jinzhu/gorm"
)

// Meeting describes a meeting between two students with the corresponding event id and date that it is happening
type Meeting struct {
	gorm.Model
	StudentOneID uint `gorm:"not null"`
	StudentTwoID uint `gorm:"not null"`
	Confirmed    bool `gorm:"not null"`
	EventID      string
	StartTime    string
	EndTime      string
}
