package model

type Group struct {
	User           User
	Article        Article
	ArticleContent ArticleContent
	Category       Category
	Tag            Tag
	ArticleTag     ArticleTag
}

var Model = new(Group)
