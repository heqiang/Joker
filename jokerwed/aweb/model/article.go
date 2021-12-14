package model

import "gorm.io/gorm"

// Article 文章表
type Article struct {
	gorm.Model
	ArticleId int64  `json:"articleid" gorm:"type:bigint;column:articleid" comment:"文章id"`
	Domain    string `json:"domain" gorm:"type:varchar(200);column:domain" comment:"域名"`
	Url       string `json:"url" gorm:"column:url" comment:"文章url"`
	Title     string `json:"title" gorm:"column:title" comment:"文章标题"`
	Content   string `json:"content" gorm:"column:content;type:text" comment:"文章内容"`
	PubTime   string `gorm:"column:pubtime" comment:"文章发布时间"`
	ClickNum  int    `json:"clicknum" gorm:"column:clicknum;default:0" comment:"文章点击数"`
	Category  string `json:"category" gorm:"column:category;default:'微段子'"`
	UserId    int64  `json:"user_id" gorm:"type:bigint;column:userid"`
}

func (article *Article) TableName() string {
	return "article"
}

// Vote 点赞
type Vote struct {
	gorm.Model
	ArticleId int64 `json:"articleid,string" comment:"文章的id"`
	Direction int   `json:"direction,string" binding:"oneof=1 0 -1" comment:"1 赞成 0 取消赞成或者反对 -1反对 "`
}

// ParamPostList 分页
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size"  form:"size"`
	Order string `json:"order" form:"order"`
}
