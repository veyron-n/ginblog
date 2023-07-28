package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

func (Category) TableName() string {
	return "category" // 返回你想要使用的表名
}
