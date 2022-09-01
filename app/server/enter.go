package server

type Group struct {
	User User
	Tag  Tag
}

var Server = new(Group)
