package actors

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	Firstname string
	Lastname string
}