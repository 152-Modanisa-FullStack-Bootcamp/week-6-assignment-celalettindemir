package handler

import (
	"net/http"
	"strings"
	"wallet/controllers"
)

type WalletHandler interface {
	WalletOperation(w http.ResponseWriter, r *http.Request)
}

type walletHandler struct {
	wc controllers.WalletController
}

func (c *walletHandler) WalletOperation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		username := strings.TrimPrefix(r.URL.Path, "/")
		if username == "" {
			c.wc.GetAllWallets(w, r)
		} else {
			c.wc.GetWalletByUsername(w, r)
		}
	} else if r.Method == http.MethodPost {

		c.wc.CreateWallet(w, r)
	} else if r.Method == http.MethodPut {
		c.wc.UpdateWallet(w, r)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

}
func NewWalletHandler(wc controllers.WalletController) WalletHandler {
	return &walletHandler{wc: wc}
}
