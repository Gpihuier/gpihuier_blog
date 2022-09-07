package controller

type Group struct {
	User    User
	Tag     Tag
	Article Article
}

var Controller = new(Group)
