package service

import (
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/model"
)

type Article interface {
	PostArticle(article model.Article) error
	UpdateArticle(article articletype.Article) error
	GetArticleByArticleId(articleId string) (model.Article, error)
	GetAllarticle(page, size int) ([]model.Article, int, error)
}
