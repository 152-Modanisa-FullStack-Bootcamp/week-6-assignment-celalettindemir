package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet/mock"
	"wallet/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllWallets(t *testing.T) {

	t.Run("only GET method allowed", func(t *testing.T) {

		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().Wallets().Return([]model.Wallet{
			{UserName: "celalettin", Amount: 100},
		}, nil).Times(1)
		wallet_Controller := NewWalletController(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		wallet_Controller.GetAllWallets(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("POST not allowed", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.GetAllWallets(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("Should return all products correctly", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))

		mockService.EXPECT().Wallets().Return([]model.Wallet{
			{UserName: "celalettin", Amount: 100},
		}, nil).Times(1)

		wallet_Controller := NewWalletController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.GetAllWallets(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", res.Result().Header.Get("content-type"))

		expectedResBody := []model.Wallet{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody[0].Amount, 100)
	})
	t.Run("should return server internal error when product service failed", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))

		mockService.EXPECT().Wallets().Return([]model.Wallet{}, errors.New("error occured")).Times(1)

		wallet_Controller := NewWalletController(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		wallet_Controller.GetAllWallets(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}
func TestGetWalletByUsername(t *testing.T) {
	t.Run("only GET method allowed", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			Wallet("celal").
			Return(model.Wallet{UserName: "celal", Amount: 0}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)
		req := httptest.NewRequest(http.MethodGet, "/celal", http.NoBody)
		res := httptest.NewRecorder()
		wallet_Controller.GetWalletByUsername(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("POST not allowed", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.GetWalletByUsername(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("Find user", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			Wallet("celal").
			Return(model.Wallet{UserName: "celal", Amount: 0}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/celal", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.GetWalletByUsername(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", res.Result().Header.Get("content-type"))

		expectedResBody := model.Wallet{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody.Amount, 0)
		assert.Equal(t, expectedResBody.UserName, "celal")
	})
	t.Run("User not found", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			Wallet("celalettin").
			Return(model.Wallet{}, errors.New("User not found")).
			Times(1)
		wallet_Controller := NewWalletController(mockService)

		req := httptest.NewRequest(http.MethodGet, "/celalettin", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.GetWalletByUsername(res, req)

		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
		assert.Equal(t, "User not found", string(res.Body.Bytes()))
	})
}
func TestPutWallet(t *testing.T) {
	t.Run("only Put method allowed", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			CreateWallet("celal").
			Return(model.Wallet{UserName: "celal", Amount: 0}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)
		req := httptest.NewRequest(http.MethodPut, "/celal", http.NoBody)
		res := httptest.NewRecorder()
		wallet_Controller.PutWallet(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("POST not allowed", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.PutWallet(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("Url username empty", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodPut, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.PutWallet(res, req)
		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
		assert.Equal(t, "Username is empty", string(res.Body.Bytes()))
	})
	t.Run("Find user", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			CreateWallet("celal").
			Return(model.Wallet{UserName: "celal", Amount: 0}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)

		req := httptest.NewRequest(http.MethodPut, "/celal", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.PutWallet(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", res.Result().Header.Get("content-type"))

		expectedResBody := model.Wallet{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody.Amount, 0)
		assert.Equal(t, expectedResBody.UserName, "celal")
	})
}
func TestPostWallet(t *testing.T) {
	t.Run("only Post method allowed", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			CashOperation("celal", 0).
			Return(model.Wallet{UserName: "celal", Amount: 0}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)

		body, _ := json.Marshal(model.Operation{Balance: 0})
		req := httptest.NewRequest(http.MethodPost, "/celal", bytes.NewReader(body))
		res := httptest.NewRecorder()
		wallet_Controller.PostWallet(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("Get not allowed", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.PostWallet(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("Url username empty", func(t *testing.T) {
		wallet_Controller := NewWalletController(nil)

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)
		res := httptest.NewRecorder()

		wallet_Controller.PostWallet(res, req)
		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
		assert.Equal(t, "Username is empty", string(res.Body.Bytes()))
	})
	t.Run("User deposit money", func(t *testing.T) {
		mockService := mock.NewMockWalletService(gomock.NewController(t))
		mockService.EXPECT().
			CashOperation("celal", 10).
			Return(model.Wallet{UserName: "celal", Amount: 10}, nil).
			Times(1)
		wallet_Controller := NewWalletController(mockService)

		body, _ := json.Marshal(model.Operation{Balance: 10})
		req := httptest.NewRequest(http.MethodPost, "/celal", bytes.NewReader(body))
		res := httptest.NewRecorder()

		wallet_Controller.PostWallet(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", res.Result().Header.Get("content-type"))

		expectedResBody := model.Wallet{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody.Amount, 10)
		assert.Equal(t, expectedResBody.UserName, "celal")
	})
}
