package service

import "jokerweb/aweb/model"

type Article interface {
	PostArticle(article model.Article) error
	UpdateArticle(article model.Article) error
	GetArticleByArticleId(articleId int64) (model.Article, error)
	GetAllarticle() []model.Article
}
