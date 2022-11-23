package rpc

type AllClient struct {
	Dorm  Dorm
	Team  Team
	Token Token
	User  User
}

var Client AllClient
