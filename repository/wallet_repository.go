package repository

import (
	"errors"
	"wallet/config"
)

var Db = map[string]int{"celal": 0, "kemal": 0}

type WalletRepository interface {
	FindAllWallet() (map[string]int, error)
	GetWalletByUserName(username string) (int, error)
	AddWallet(username string) (int, error)
	CashOperation(username string, Amount int) (int, error)
}

type walletRepository struct {
}

func (r *walletRepository) FindAllWallet() (map[string]int, error) {
	return Db, nil
}

func (r *walletRepository) GetWalletByUserName(username string) (int, error) {
	v, ok := Db[username]
	if !ok {
		return 0, errors.New("Not Find UserName")
	}
	return v, nil
}
func (r *walletRepository) AddWallet(username string) (int, error) {
	v, ok := Db[username]
	if !ok {
		Db[username] = config.C.InitialBalanceAmount
		return Db[username], nil
	}
	return v, nil
}
func (r *walletRepository) CashOperation(username string, Amount int) (int, error) {
	_, ok := Db[username]
	if !ok {
		return 0, errors.New("Not Find UserName")
	}
	if Amount < 0 {
		if Db[username]+Amount < config.C.MinumumBalanceAmount {
			return 0, errors.New("Cash can't be withdrawn")
		}
	}
	Db[username] += Amount
	return Db[username], nil
}
func NewWalletRepository() WalletRepository {
	return &walletRepository{}
}
