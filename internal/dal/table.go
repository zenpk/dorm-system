package dal

type AllTable struct {
	Account  Account
	Building Building
	Dorm     Dorm
	Order    Order
	Team     Team
	TeamUser TeamUser
	Token    Token
	User     User
}

var Table AllTable
