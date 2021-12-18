package service

import (
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/model"
)

type Article interface {
	PostArticle(article model.Article) error
	UpdateArticle(article articletype.Article) error
	GetArticleByArticleId(articleId string) (model.Article, error)
	GetAllArticle(page, size int) ([]model.Article, int, error)
	VoteArticle(vote model.Vote, userId int64) error
}
