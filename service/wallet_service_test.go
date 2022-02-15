package service

import (
	"testing"
	"wallet/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWallets(t *testing.T) {
	t.Run("Get All Wallets", func(t *testing.T) {
		mockData := map[string]int{
			"celalettin": 100,
			"celal":      10,
		}
		mockService := mock.NewMockWalletRepository(gomock.NewController(t))

		mockService.EXPECT().FindAllWallet().Return(mockData, nil).Times(1)
		wallet_Service := NewWalletService(mockService)

		wallets, _ := wallet_Service.Wallets()
		for _, w := range wallets {
			assert.Equal(t, mockData[w.UserName], w.Amount)
		}
	})
}

func TestWallet(t *testing.T) {
	t.Run("Get Wallet Amount", func(t *testing.T) {
		mockService := mock.NewMockWalletRepository(gomock.NewController(t))

		mockService.EXPECT().GetWalletByUserName("celal").Return(0, nil).Times(1)
		wallet_Service := NewWalletService(mockService)

		wallet, _ := wallet_Service.Wallet("celal")
		assert.Equal(t, wallet.Amount, 0)
	})
}

func TestCreateWallet(t *testing.T) {
	t.Run("New user amount", func(t *testing.T) {
		mockService := mock.NewMockWalletRepository(gomock.NewController(t))

		mockService.EXPECT().
			AddWallet("celalettin").
			Return(100, nil).
			Times(1)
		wallet_Service := NewWalletService(mockService)

		wallet, _ := wallet_Service.CreateWallet("celalettin")
		assert.Equal(t, wallet.Amount, 100)
	})
}
func TestCashOperation(t *testing.T) {
	t.Run("Get Wallet Amount", func(t *testing.T) {
		mockService := mock.NewMockWalletRepository(gomock.NewController(t))

		mockService.EXPECT().CashOperation("celalettin", 100).Return(0, nil).Times(1)
		wallet_Service := NewWalletService(mockService)

		wallet, _ := wallet_Service.CashOperation("celalettin", 100)
		assert.Equal(t, wallet.Amount, 0)
	})
}
