package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"spider/global"
	"strconv"
)

type Article struct {
	gorm.Model
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid"`
	Domain    string `json:"domain" gorm:"type:varchar(200);column:domain"`
	Url       string `json:"url" gorm:"column:url"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content;type:text"`
	PubTime   string `gorm:"column:pubtime"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0"`
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
}

type EsItem struct {
	gorm.Model
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid"`
	Domain    string `json:"domain" gorm:"type:varchar(200);column:domain"`
	Url       string `json:"url" gorm:"column:url"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content;type:text"`
	PubTime   string `gorm:"column:pubtime"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0"`
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
}

func (item *Article) BeforeCreate(tx *gorm.DB) error {
	esData := EsItem{
		ArticleId: item.ArticleId,
		Domain:    item.Domain,
		Url:       item.Url,
		Title:     item.Title,
		Content:   item.Content,
		PubTime:   item.PubTime,
		ClickNum:  item.ClickNum,
		Category:  item.Category,
	}

	esRes, err := global.Es.Index().Index("joker").Id(strconv.FormatInt(item.ArticleId, 10)).BodyJson(esData).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(esRes.Id)
	return nil
}
