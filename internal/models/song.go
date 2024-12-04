package models

import (
	"github.com/jackc/pgx/pgtype"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Group       string           `gorm:"type:varchar(255);not null" json:"group"`
	Song        string           `gorm:"type:varchar(255);not null" json:"song"`
	ReleaseDate string           `gorm:"type:date;not null" json:"releaseDate"`
	Verses      pgtype.TextArray `gorm:"type:text[];not null" json:"verses"`
	Link        string           `gorm:"type:varchar(255);not null" json:"link"`
}
