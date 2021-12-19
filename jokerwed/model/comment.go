package model

import "gorm.io/gorm"

type Comment struct {
	gorm.DB
	CommentId           int64  `json:"commentid" gorm:"type:bigint;column:commentid;comment:'文章id'"`
	UserId              int64  `json:"userid" gorm:"type:bigint;column:userid;comment:'评论人的id'"`
	ArticleId           int64  `json:"articleid" gorm:"type:bigint;column:articleid;comment:'被评论的文章id'"`
	ParentCommentId     int64  `json:"parentcommentid" gorm:"type:bigint;column:parentcommentid;comment:'父级评论id'"`
	ParentCommentUserId int64  `json:"parentcommentuserid" gorm:"type:bigint;column:parentcommentuserid;comment:'父级评论的用户id'"`
	ReplayCommentId     int64  `json:"replaycommentid" gorm:"type:bigint;column:replaycommentid;comment:'被回复的评论id'"`
	ReplayCommentUserId int64  `json:"replaycommentuserid" gorm:"type:bigint;column:replaycommentuserid;comment:'被回复的评论的用户id'"`
	CommentLevel        int    `json:"commentlevel" gorm:"type:int;column:'commentlevel';comment:'评论级别 文章1 其余均为2'"`
	Content             string `json:"content" gorm:"type:varchar(255);column:content;comment:'评论内容'"`
}
