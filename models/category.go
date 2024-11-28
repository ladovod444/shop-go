package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Title    string    `json:"Title" example:"Category title"`
	Slug     string    `json:"Slug" example:"Category slug"`
	Products []Product `json:"Products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
