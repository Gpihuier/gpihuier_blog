package route

type Group struct {
	UserRouters UserRouters
	TagRouters  TagRouters
}

var RouterEnter = new(Group)
