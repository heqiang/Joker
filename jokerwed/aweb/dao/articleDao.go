package dao

import (
	"errors"
	"jokerweb/aweb/model"
	"jokerweb/global"
	"jokerweb/utils/jwt"
	"jokerweb/utils/snowflake"
	"time"
)

const (
	ArticleNotExist = "该文章不存在"
	ArticleExist    = "文章已存在"
)

func PostArticle(article model.Article) error {
	_, err := GetArticleByArticleId(article.ArticleId)
	if err != nil {
		art := model.Article{
			ArticleId: snowflake.GetSnowId(),
			Domain:    article.Domain,
			Url:       article.Url,
			Content:   article.Content,
			PubTime:   time.Now().Format("2006-01-02 15:04:05"),
			Category : article.Category,
			UserId:
		}
		global.Db.Create(&art)
	}
	return errors.New(ArticleExist)
}
func UpdateArticle(article model.Article) error {
	art := model.Article{
		Title:   article.Title,
		Content: article.Content,
	}
	res := global.Db.Create(&art)
	if res.RowsAffected == 0 {
		return res.Error
	}
	return nil

}
func GetArticleByArticleId(articleId int64) (model.Article, error) {
	var article model.Article
	res := global.Db.Where("articleid=?", articleId).First(&article)
	if res.RowsAffected == 0 {
		return article, errors.New(ArticleNotExist)
	}
	return article, nil
}
func GetAllarticle() (allArticle []model.Article) {
	allArticle = []model.Article{}
	global.Db.Find(&allArticle)
	return
}
