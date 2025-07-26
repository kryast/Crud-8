package models

import "gorm.io/gorm"

type Employee struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Position  string         `json:"position"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
