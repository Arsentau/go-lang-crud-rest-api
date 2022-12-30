package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id          string `gorm:"not null; unique_index"`
	Title       string `gorm:"not null"`
	Bought      bool   `gorm:"default:false"`
	Description string
}

// type Director struct {
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// }
