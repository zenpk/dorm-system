package dal

type TableSet struct {
	Building       *Building
	Dorm           *Dorm
	Order          *Order
	UserCredential *UserCredential
	UserInfo       *UserInfo
}

var Table TableSet
