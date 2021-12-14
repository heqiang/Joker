package implements

import (
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/aweb/dao"
	"jokerweb/model"
)

type Article struct{}

func (A *Article) PostArticle(article model.Article) error {
	return dao.PostArticle(article)
}
func (A *Article) UpdateArticle(article articletype.Article) error {
	return dao.UpdateArticle(article)
}
func (A *Article) GetArticleByArticleId(articleId int64) (model.Article, error) {
	return dao.GetArticleByArticleId(articleId)
}

func (A *Article) GetAllarticle() []model.Article {
	return dao.GetAllarticle()
}
