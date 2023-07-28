package model

import "github.com/jinzhu/gorm"

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type: varchar(20); not null " json: "title"`
	Cid     int    `gorm:"type: int; not null " json: "cid"`
	Desc    string `gorm:"type: varchar(200); not null " json: "desc"`
	Context string `gorm:"type: longtext; not null " json: "context"`
	Img     string `gorm:"type: varchar(100);" json: "img"`
}
