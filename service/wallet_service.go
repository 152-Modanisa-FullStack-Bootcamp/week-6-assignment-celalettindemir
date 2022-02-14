package service

import (
	"wallet/model"
	"wallet/repository"
)

type WalletService interface {
	Wallets() ([]model.Wallet, error)
	Wallet(username string) (model.Wallet, error)
	CreateWallet(username string) (model.Wallet, error)
	CashOperation(username string, Amount int) (model.Wallet, error)
}

type walletService struct {
	repo repository.WalletRepository
}

func (c *walletService) Wallets() ([]model.Wallet, error) {
	wallets, _ := c.repo.FindAllWallet()
	retWallets := []model.Wallet{}
	for i, w := range wallets {
		retWallets = append(retWallets, model.Wallet{UserName: i, Amount: w})
	}
	return retWallets, nil
}
func (c *walletService) Wallet(username string) (model.Wallet, error) {
	amount, err := c.repo.GetWalletByUserName(username)
	if err != nil {
		return model.Wallet{}, err
	}
	retWallets := model.Wallet{UserName: username, Amount: amount}
	return retWallets, nil
}
func (c *walletService) CreateWallet(username string) (model.Wallet, error) {
	amount, err := c.repo.AddWallet(username)
	if err != nil {
		return model.Wallet{}, err
	}
	retWallets := model.Wallet{UserName: username, Amount: amount}
	return retWallets, nil
}
func (c *walletService) CashOperation(username string, Amount int) (model.Wallet, error) {
	amount, err := c.repo.CashOperation(username, Amount)
	if err != nil {
		return model.Wallet{}, err
	}
	retWallets := model.Wallet{UserName: username, Amount: amount}
	return retWallets, nil
}
func NewWalletService(repo repository.WalletRepository) WalletService {
	return &walletService{repo: repo}
}
