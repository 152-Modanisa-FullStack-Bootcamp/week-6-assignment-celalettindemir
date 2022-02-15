package model

type Wallet struct {
	UserName string `json:"username"`
	Amount   int    `json:"amount"`
}
type Operation struct {
	Balance int `json:"balance"`
}
