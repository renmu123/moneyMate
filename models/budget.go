package models

import "github.com/jinzhu/gorm"

type Budget struct {
	gorm.Model
	Name  string `validate:"required" json:"name"`
	Price float64
}
