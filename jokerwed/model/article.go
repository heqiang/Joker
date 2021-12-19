package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"jokerweb/global"
	"strconv"
)

// Article 文章表
type Article struct {
	gorm.Model
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid;comment:'文章id'" `
	Title     string `json:"title" gorm:"column:title;comment:'文章标题'"  binding:"required"`
	Content   string `json:"content" gorm:"column:content;type:text; comment:'文章内容'" binding:"required"`
	PubTime   string `json:"pub_time" gorm:"column:pubtime;comment:'文章发布时间'"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0;comment:'文章点击量'" `
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
	UserId    int64  `json:"user_id" gorm:"type:bigint;column:userid;comment:'用户id'"`
	VoteNum   int    `json:"votenum" gorm:"type:int;column:votenum;default 0;comment:'文章点赞量'"`
}

func (item *Article) TableName() string {
	return "article"
}

// Vote 点赞
type Vote struct {
	ArticleId int64   `json:"articleid,string" comment:"文章的id" binging:"required"`
	Direction float64 `json:"direction,string" comment:"1赞成 0 取消点赞" binging:"required"`
}

// ParamPostList 分页
type ParamPostList struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data" form:"data"`
}

type EsItem struct {
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content;type:text"`
	PubTime   string `gorm:"column:pubtime"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0"`
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
}

func (item *Article) AfterCreate(tx *gorm.DB) error {
	esData := EsItem{
		ArticleId: item.ArticleId,
		Title:     item.Title,
		Content:   item.Content,
		PubTime:   item.PubTime,
		ClickNum:  item.ClickNum,
		Category:  item.Category,
	}

	res, err := global.Es.Index().Index("joker").Id(strconv.FormatInt(item.ArticleId, 10)).BodyJson(esData).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(res.Id)
	return nil
}
