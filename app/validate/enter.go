package validate

type Group struct {
	User    User
	Tag     Tag
	Article Article
}

var Validate = new(Group)
