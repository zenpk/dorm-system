package rpc

type AllClient struct {
	Dorm  Dorm
	Order Order
	Team  Team
	Token Token
	User  User
}

var Client AllClient
