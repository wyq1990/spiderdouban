package service

import (
	"douban/model"
	"fmt"
)

// CreateVideoService 视频投稿的服务
type InsertMovieService struct {
	DoubanId  int    `form:"doubanId" json:"doubanId" binding:"required`
	MovieName string `form:"movieName" json:"movieName" binding:"max=3000"`
	// URL    string `form:"url" json:"url"`
	// Avatar string `form:"avatar" json:"avatar"`
	MovieContent string `form:"movieContent" json:"movieContent"`
	MovieImg     string `form:"movieImg" json:"movieImg"`
}

func (service *InsertMovieService) Insert() {
	movie := model.Movie{
		DoubanId:     service.DoubanId,
		MovieName:    service.MovieName,
		MovieContent: service.MovieContent,
		MovieImg:     service.MovieImg,
	}

	err := model.DB.Create(&movie).Error
	if err != nil {
		fmt.Println("插入失败")
	} else {
		fmt.Println("插入成功")
	}
}
