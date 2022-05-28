package models

import (
	"gorm.io/gorm"
)

type Actress struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Age  uint
	Name string
}
