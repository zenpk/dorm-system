package dal

type TableSet struct {
	Building       *Building
	Dorm           *Dorm
	Order          *Order
	Team           *Team
	UserCredential *Account
	UserInfo       *User
}

var Table TableSet
