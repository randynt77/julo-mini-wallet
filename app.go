package main

import (
	"julo-mini-wallet/config"
	wallet_handler "julo-mini-wallet/wallet/handler"
	wallet_repo "julo-mini-wallet/wallet/repo"
	wallet_usecase "julo-mini-wallet/wallet/usecase"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"
)

func main() {
	// Init Config
	dbConnection := config.GetDatabaseConns()
	router := httprouter.New()

	//Init Repository
	repoRepo := wallet_repo.NewRepo(dbConnection.Master)

	//init usecase
	usecase := wallet_usecase.New(repoRepo, dbConnection.Master)

	//Init Handler router
	wallet_handler.RegisterRoute(router, usecase)

	// start gracefull service
	err := grace.Serve(":8080", router)
	if err != nil {
		return
	}
}
