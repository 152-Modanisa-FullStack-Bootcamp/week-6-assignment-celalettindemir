package repository

import (
	"testing"
	"wallet/config"

	"github.com/stretchr/testify/assert"
)

func TestGetWalletByUserName(t *testing.T) {
	t.Run("Dont Exist User", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		_, err := wallet_repository.GetWalletByUserName("celalettin")
		assert.Equal(t, err.Error(), "Not Find UserName")
	})
	t.Run("Exist User", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		walletAmount, _ := wallet_repository.GetWalletByUserName("celal")
		assert.Equal(t, walletAmount, 0)
	})
}
func TestAddWallet(t *testing.T) {

	t.Run("Dont Exist User return new", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		walletAmount, _ := wallet_repository.AddWallet("celalettin")
		assert.Equal(t, walletAmount, config.C.InitialBalanceAmount)
	})
	t.Run("Exist User return same", func(t *testing.T) {
		Db = map[string]int{"celal": 10}
		wallet_repository := NewWalletRepository()
		walletAmount, _ := wallet_repository.AddWallet("celal")
		assert.Equal(t, walletAmount, 10)

	})
}
func TestCashOperation(t *testing.T) {

	t.Run("Dont Exist User", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		_, err := wallet_repository.CashOperation("celalettin", 10)
		assert.Equal(t, err.Error(), "Not Find UserName")
	})
	t.Run("Havent enough money", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		_, err := wallet_repository.CashOperation("celal", -150)
		assert.Equal(t, err.Error(), "Cash can't be withdrawn")
	})
	t.Run("Exist User have enough", func(t *testing.T) {
		Db = map[string]int{"celal": 0}
		wallet_repository := NewWalletRepository()
		wallet, _ := wallet_repository.CashOperation("celal", -10)
		assert.Equal(t, wallet, -10)
	})
}
