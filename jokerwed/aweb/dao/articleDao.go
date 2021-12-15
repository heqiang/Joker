package dao

import (
	"errors"
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/global"
	"jokerweb/model"
	"jokerweb/utils/snowflake"
	"time"
)

var ArticleNotExist = errors.New("文章不存在")

func PostArticle(article model.Article) error {
	art := model.Article{
		ArticleId: snowflake.GetSnowId(),
		Domain:    article.Domain,
		Url:       article.Url,
		Title:     article.Title,
		Content:   article.Content,
		PubTime:   time.Now().Format("2006-01-02 15:04:05"),
		Category:  article.Category,
		UserId:    article.UserId,
	}
	tx := global.Db.Begin()
	tx.Create(&art)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()

	return nil
}
func UpdateArticle(article articletype.Article) error {
	tx := global.Db.Begin()
	tx.Where("articleId=?", article.ArticleId).Save(&article)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil

}
func GetArticleByArticleId(articleId string) (model.Article, error) {
	var art model.Article
	res := global.Db.Where("articleid=?", articleId).First(&art)
	if res.RowsAffected == 0 {
		return art, ArticleNotExist
	}
	return art, nil
}
func GetAllarticle() (allArticle []model.Article) {
	allArticle = []model.Article{}
	global.Db.Find(&allArticle)
	return
}
