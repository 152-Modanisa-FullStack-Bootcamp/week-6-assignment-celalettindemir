package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"wallet/service"
)

type WalletController interface {
	GetAllWallets(w http.ResponseWriter, r *http.Request)
	WalletOperation(w http.ResponseWriter, r *http.Request)
}

type walletController struct {
	wsvc service.WalletService
}

type Operation struct {
	Balance string `json:"balance"`
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
func (c *walletController) WalletOperation(w http.ResponseWriter, r *http.Request) {

	username := strings.TrimPrefix(r.URL.Path, "/")
	if r.Method == http.MethodGet {
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
	} else if r.Method == http.MethodPost {
		var o Operation
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if r.Method == http.MethodPut {
		username := strings.TrimPrefix(r.URL.Path, "/")
		wallet, err := c.wsvc.CreateWallet(username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		walletsBytes, _ := json.Marshal(wallet)
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		w.Write(walletsBytes)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

}
func NewWalletController(wsvc service.WalletService) WalletController {
	return &walletController{wsvc: wsvc}
}
