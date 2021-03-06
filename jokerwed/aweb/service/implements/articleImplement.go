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
func (A *Article) GetArticleByArticleId(articleId string) (model.Article, error) {
	return dao.GetArticleByArticleId(articleId)
}

func (A *Article) GetAllArticle(page, size int) ([]model.Article, int, error) {
	return dao.GetAllarticle(page, size)
}

func (A *Article) VoteArticle(vote model.Vote, userId int64) error {
	return dao.VoteArticle(vote, userId)
}

func (A *Article) CommentArticle(comment articletype.CommentArticleParam, userId int64) error {
	return dao.CommentArticle(comment, userId)
}

func (A *Article) CommentToComment(comment articletype.CommentToCommentParam, userId int64) error {

	return dao.CommentToComment(comment, userId)

}
