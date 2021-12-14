package service

import (
	"jokerweb/aweb/controller/article"
	"jokerweb/model"
)

type Article interface {
	PostArticle(article model.Article) error
	UpdateArticle(article article.Article) error
	GetArticleByArticleId(articleId int64) (model.Article, error)
	GetAllarticle() []model.Article
}
