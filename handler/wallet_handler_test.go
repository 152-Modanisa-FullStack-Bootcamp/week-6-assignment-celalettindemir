package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet/mock"
	"wallet/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWalletOperation(t *testing.T) {

	t.Run("Get All Wallets", func(t *testing.T) {
		mockService := mock.NewMockWalletController(gomock.NewController(t))
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		res.WriteHeader(http.StatusOK)
		mockService.EXPECT().
			GetAllWallets(res, req).
			Return().
			Times(1)
		wallet_Handler := NewWalletHandler(mockService)

		wallet_Handler.WalletOperation(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("Get user Wallet", func(t *testing.T) {
		mockService := mock.NewMockWalletController(gomock.NewController(t))
		req := httptest.NewRequest(http.MethodGet, "/celal", http.NoBody)
		res := httptest.NewRecorder()
		res.WriteHeader(http.StatusOK)
		mockService.EXPECT().
			GetWalletByUsername(res, req).
			Return().
			Times(1)
		wallet_Handler := NewWalletHandler(mockService)

		wallet_Handler.WalletOperation(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("Put create user wallet ", func(t *testing.T) {
		mockService := mock.NewMockWalletController(gomock.NewController(t))
		req := httptest.NewRequest(http.MethodPut, "/celal", http.NoBody)
		res := httptest.NewRecorder()
		res.WriteHeader(http.StatusOK)
		mockService.EXPECT().
			PutWallet(res, req).
			Return().
			Times(1)
		wallet_Handler := NewWalletHandler(mockService)

		wallet_Handler.WalletOperation(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("Post Cash Operation", func(t *testing.T) {
		mockService := mock.NewMockWalletController(gomock.NewController(t))
		body, _ := json.Marshal(model.Operation{Balance: 10})
		req := httptest.NewRequest(http.MethodPost, "/celal", bytes.NewReader(body))
		res := httptest.NewRecorder()
		res.WriteHeader(http.StatusOK)
		mockService.EXPECT().
			PostWallet(res, req).
			Return().
			Times(1)
		wallet_Handler := NewWalletHandler(mockService)

		wallet_Handler.WalletOperation(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("Invalid Request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/celal", http.NoBody)
		res := httptest.NewRecorder()
		wallet_Handler := NewWalletHandler(nil)

		wallet_Handler.WalletOperation(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
}
