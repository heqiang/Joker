package articletype

import (
	"context"
	"gorm.io/gorm"
	"jokerweb/global"
	"jokerweb/model"
	"strconv"
)

type Article struct {
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid" comment:"文章id" binding:"required"`
	Domain    string `json:"domain" gorm:"type:varchar(200);column:domain" comment:"域名"`
	Url       string `json:"url" gorm:"column:url" comment:"文章url"`
	Title     string `json:"title" gorm:"column:title" comment:"文章标题" `
	Content   string `json:"content" gorm:"column:content;type:text" comment:"文章内容"`
	PubTime   string `gorm:"column:pubtime" comment:"文章发布时间"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0" comment:"文章点击数"`
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
}

func (item *Article) AfterUpdate(tx *gorm.DB) (err error) {
	_, err = global.Es.Delete().Index("joker").Id(strconv.FormatInt(item.ArticleId, 10)).Do(context.Background())
	if err != nil {
		return err
	}
	var articleInfo Article
	res := global.Db.Where("articleId=?", item.ArticleId).Take(&articleInfo)
	if res.RowsAffected == 0 {
		return res.Error
	}
	esData := model.EsItem{
		ArticleId: item.ArticleId,
		Title:     item.Title,
		Content:   item.Content,
		PubTime:   articleInfo.PubTime,
		ClickNum:  articleInfo.ClickNum,
		Category:  articleInfo.Category,
	}

	_, err = global.Es.Index().Index("joker").Id(strconv.FormatInt(item.ArticleId, 10)).BodyJson(esData).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

type ArticlParam struct {
	Title   string `json:"title" gorm:"column:title" comment:"文章标题" binding:"required" `
	Content string `json:"content" gorm:"column:content;type:text" comment:"文章内容" binding:"required" `
}
