package main

import (
	"fmt"
	"net/http"
	"wallet/controllers"
	"wallet/repository"
	"wallet/service"
)

func main() {
	walletRepo := repository.NewWalletRepository()
	walletService := service.NewWalletService(walletRepo)
	walletController := controllers.NewWalletController(walletService)

	http.HandleFunc("/:username", walletController.WalletOperation)
	http.HandleFunc("/", walletController.GetAllWallets)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
