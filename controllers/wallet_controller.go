package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"wallet/model"
	"wallet/service"
)

type WalletController interface {
	GetAllWallets(w http.ResponseWriter, r *http.Request)
	GetWalletByUsername(w http.ResponseWriter, r *http.Request)
	PutWallet(w http.ResponseWriter, r *http.Request)
	PostWallet(w http.ResponseWriter, r *http.Request)
}

type walletController struct {
	wsvc service.WalletService
}

func (c *walletController) GetAllWallets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	wallets, err := c.wsvc.Wallets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	walletsBytes, _ := json.Marshal(wallets)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(walletsBytes)
	w.WriteHeader(http.StatusOK)
}
func (c *walletController) GetWalletByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	username := strings.TrimPrefix(r.URL.Path, "/")
	wallet, err := c.wsvc.Wallet(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	walletsBytes, _ := json.Marshal(wallet)
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(walletsBytes)
	w.WriteHeader(http.StatusOK)
}
func (c *walletController) PutWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	username := strings.TrimPrefix(r.URL.Path, "/")
	if username == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Username is empty"))
		return
	}
	wallet, err := c.wsvc.CreateWallet(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	walletsBytes, _ := json.Marshal(wallet)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)
}
func (c *walletController) PostWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	username := strings.TrimPrefix(r.URL.Path, "/")
	if username == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Username is empty"))
		return
	}
	var o model.Operation
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wallet, err := c.wsvc.CashOperation(username, o.Balance)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	walletsBytes, _ := json.Marshal(wallet)
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(walletsBytes)
}
func NewWalletController(wsvc service.WalletService) WalletController {
	return &walletController{wsvc: wsvc}
}
