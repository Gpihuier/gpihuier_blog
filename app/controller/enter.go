package controller

type Group struct {
	User User
	Tag  Tag
}

var Controller = new(Group)
