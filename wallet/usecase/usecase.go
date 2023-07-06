package usecase

import (
	"julo-mini-wallet/util"
	"julo-mini-wallet/wallet"

	"github.com/jmoiron/sqlx"
)

type (
	usecase struct {
		walletRepository wallet.Repository
		pqConn           *sqlx.DB
	}
)

func New(walletRepository wallet.Repository, dbConn *sqlx.DB) *usecase {
	return &usecase{
		walletRepository: walletRepository,
		pqConn:           dbConn,
	}
}

func (u *usecase) CreateAccount(userID string) (userToken wallet.NewAccountResponse, err error) {

	userToken = wallet.NewAccountResponse{
		Token: util.GenerateJWTToken(userID),
	}
	_, getErr := u.walletRepository.GetAccount(userID)
	if getErr == nil {
		return
	}
	err = u.walletRepository.CreateAccount(userID)
	if err != nil {
		return
	}

	return
}
