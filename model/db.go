package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败,请检查参数: ", err)
	}

	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// 禁止默认表名的复数形式
	db.SingularTable(true)

	// SetMaxIdleConns 设置空闲连接池的最大连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置数据库打开连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
