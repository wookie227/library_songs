package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group       string `gorm:"type:varchar(255);not null" json:"group"`
	Song        string `gorm:"type:varchar(255);not null" json:"song"`
	ReleaseDate string `gorm:"type:date;not null" json:"releaseDate"`
	Text        string `gorm:"type:text;not null" json:"text"`
	Link        string `gorm:"type:varchar(255);not null" json:"link"`
}
