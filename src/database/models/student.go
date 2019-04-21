package models

import "github.com/jinzhu/gorm"

// Student struct
type Student struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Consent     bool   `gorm:"default:false"`
	Preferences string
	Location    string
}
