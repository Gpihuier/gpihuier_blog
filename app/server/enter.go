package server

type Group struct {
	User    User
	Tag     Tag
	Article Article
}

var Server = new(Group)
