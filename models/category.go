package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `validate:"required" json:"name"`
	Icno string
}
