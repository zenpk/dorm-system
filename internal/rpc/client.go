package rpc

type ClientSet struct {
	Dorm  Dorm
	Token Token
	User  User
}

var Client ClientSet
