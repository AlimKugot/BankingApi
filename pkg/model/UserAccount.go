package main

type UserAccount struct {
	Id       int     `json:"-"`
	UserName string  `json:"user-name"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

//
//var accounts = []UserAccount{
//	{1, "alim", "992001"},
//	{2, "pavel", "123456"},
//	{3, "anton", "random"},
//}
