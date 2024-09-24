package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName     string
	MiddleInitial string
	LastName      string
}
