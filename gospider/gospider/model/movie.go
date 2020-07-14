package model

import "github.com/jinzhu/gorm"

// Movie 电影模型
type Movie struct {
	gorm.Model
	DoubanId  int `gorm:"unique"`
	MovieName string
	MovieContent string
	MovieImg string
}
