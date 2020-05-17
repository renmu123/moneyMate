package models

import "github.com/jinzhu/gorm"

type Record struct {
	gorm.Model
	Name     string   `validate:"required" json:"name"`
	Type     int      `validate:"required,oneof=0 1" json:"type"`
	Money    float64  `validate:"required" json:"money"`
	Desc     string   `validate:"required" json:"desc"`
	Category Category `validate:"-"`
	TagIds   []Tags   `gorm:"many2many:record_tags;" json:"tag_ids"`
}
