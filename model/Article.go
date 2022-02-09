package model

import "gorm.io/gorm"

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Img     string `gorm:"type:varchar(100);not null" json:"img"`
}
