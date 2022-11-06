package dal

type TableSet struct {
	Building       *Building
	Dorm           *Dorm
	Order          *Order
	Team           *Team
	UserCredential *UserCredential
	UserInfo       *UserInfo
}

var Table TableSet
