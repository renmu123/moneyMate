package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	RecordID int    `json:"record_id" validate:"required"`
	Name     string `validate:"required" json:"name"`
	Icon     string
}
