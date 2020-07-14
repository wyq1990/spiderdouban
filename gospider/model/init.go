package model

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(sqlPath string) {
	db, err := gorm.Open("mysql", sqlPath)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	// if gin.Mode() == "release" {
	// 	db.LogMode(false)
	// }
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}