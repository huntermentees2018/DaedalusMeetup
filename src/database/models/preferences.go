package models

import "github.com/jinzhu/gorm"

// Preference struct
type Preference struct {
	gorm.Model
	DayOfWeek string
	Time      string
}
