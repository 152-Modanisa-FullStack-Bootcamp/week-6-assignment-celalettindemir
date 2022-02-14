package main

import (
	"fmt"
	"net/http"
	"wallet/controllers"
	"wallet/handler"
	"wallet/repository"
	"wallet/service"
)

func main() {
	walletRepo := repository.NewWalletRepository()
	walletService := service.NewWalletService(walletRepo)
	walletController := controllers.NewWalletController(walletService)
	walletHandler := handler.NewWalletHandler(walletController)

	http.HandleFunc("/", walletHandler.WalletOperation)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
