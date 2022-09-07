package route

type Group struct {
	UserRouters    UserRouters
	TagRouters     TagRouters
	ArticleRouters ArticleRouters
}

var RouterEnter = new(Group)
