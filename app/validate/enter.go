package validate

type Group struct {
	Common  Common
	User    User
	Tag     Tag
	Article Article
}

var Validate = new(Group)
