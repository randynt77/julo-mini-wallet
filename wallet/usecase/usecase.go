package usecase

import (
	"errors"
	"julo-mini-wallet/util"
	"julo-mini-wallet/wallet"
	"time"

	"github.com/google/uuid"
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

func (u *usecase) Withdraw(userID string, amount float64, referenceID string) (withdrawData wallet.WithDrawal, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return withdrawData, errors.New("Wallet not exist")
	}
	if walletAcount.Status == wallet.StatusDisabled {
		return withdrawData, errors.New("Wallet is disabled")
	}

	if walletAcount.Balance-amount < 0 {
		return withdrawData, errors.New("Insufficient money balance")
	}

	transactionInput := wallet.Transaction{
		WalletID:         walletAcount.ID,
		ActorID:          userID,
		TransactedAt:     time.Now(),
		Type:             wallet.TypeWithdrawal,
		Amount:           amount,
		Status:           "Success",
		InputReferenceID: referenceID,
		ReferenceID:      uuid.New().String(),
	}

	transactionID, err := u.walletRepository.InsertTransactionData(transactionInput)
	if err != nil {
		return
	}

	finalBalance := walletAcount.Balance - amount
	err = u.walletRepository.UpdateAccountBalance(userID, finalBalance)
	if err != nil {
		return
	}

	withdrawData = wallet.WithDrawal{
		ID:          transactionID,
		WithdrawnBy: userID,
		Status:      "Success",
		WithdrawnAt: transactionInput.TransactedAt,
		Amount:      amount,
		ReferenceID: transactionInput.ReferenceID,
	}

	return
}

func (u *usecase) Deposit(userID string, amount float64, referenceID string) (depositData wallet.Deposit, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return depositData, errors.New("Wallet not exist")
	}
	if walletAcount.Status == wallet.StatusDisabled {
		return depositData, errors.New("Wallet is disabled")
	}

	transactionInput := wallet.Transaction{
		WalletID:         walletAcount.ID,
		ActorID:          userID,
		TransactedAt:     time.Now(),
		Type:             wallet.TypeDeposit,
		Amount:           amount,
		Status:           "Success",
		InputReferenceID: referenceID,
		ReferenceID:      uuid.New().String(),
	}

	transactionID, err := u.walletRepository.InsertTransactionData(transactionInput)
	if err != nil {
		return
	}

	finalBalance := walletAcount.Balance + amount
	err = u.walletRepository.UpdateAccountBalance(userID, finalBalance)
	if err != nil {
		return
	}

	depositData = wallet.Deposit{
		ID:          transactionID,
		DepositedBy: userID,
		Status:      "Success",
		DepositedAt: transactionInput.TransactedAt,
		Amount:      amount,
		ReferenceID: transactionInput.ReferenceID,
	}

	return
}

func (u *usecase) GetWalletTransactions(userID string) (transactionsData []wallet.TransactionResponse, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return transactionsData, errors.New("Wallet not exist")
	}

	if walletAcount.Status == wallet.StatusDisabled {
		return transactionsData, errors.New("Wallet disabled")
	}

	transactionsData, err = u.walletRepository.GetTransactionData(walletAcount.ID)

	return
}

func (u *usecase) GetWalletBalance(userID string) (walletData wallet.EnableWalletResponse, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return walletData, errors.New("Wallet not exist")
	}

	if walletAcount.Status == wallet.StatusDisabled {
		return walletData, errors.New("Wallet disabled")
	}

	walletData = wallet.EnableWalletResponse{
		ID:        walletAcount.ID,
		OwnedBy:   walletAcount.OwnedBy,
		Status:    walletAcount.Status,
		EnabledAt: walletAcount.EnabledAt,
		Balance:   walletAcount.Balance,
	}

	return
}

func (u *usecase) DisableWallet(userID string) (walletData wallet.DisableWalletResponse, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return walletData, errors.New("Wallet not exist")
	}
	if walletAcount.Status == wallet.StatusDisabled {
		return walletData, errors.New("Already disabled")
	}
	walletInfo, err := u.walletRepository.UpdateAccountStatus(userID, wallet.StatusDisabled)
	if err != nil {
		return
	}

	walletData = wallet.DisableWalletResponse{
		ID:         walletInfo.ID,
		OwnedBy:    walletInfo.OwnedBy,
		Status:     walletInfo.Status,
		DisabledAt: walletInfo.DisabledAt,
		Balance:    walletInfo.Balance,
	}

	return
}

func (u *usecase) EnableWallet(userID string) (walletData wallet.EnableWalletResponse, err error) {
	walletAcount, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return walletData, errors.New("Wallet not exist")
	}

	if walletAcount.Status == wallet.StatusEnabled {
		return walletData, errors.New("Already enabled")
	}
	walletInfo, err := u.walletRepository.UpdateAccountStatus(userID, wallet.StatusEnabled)
	if err != nil {
		return
	}

	walletData = wallet.EnableWalletResponse{
		ID:        walletInfo.ID,
		OwnedBy:   walletInfo.OwnedBy,
		Status:    walletInfo.Status,
		EnabledAt: walletInfo.EnabledAt,
		Balance:   walletInfo.Balance,
	}

	return
}

func (u *usecase) CreateAccount(userID string) (userToken wallet.NewAccountResponse, err error) {

	userToken = wallet.NewAccountResponse{
		Token: util.GenerateJWTToken(userID),
	}
	_, getErr := u.walletRepository.GetAccount(userID)
	if getErr != nil {
		return userToken, getErr
	}

	err = u.walletRepository.CreateAccount(userID)
	if err != nil {
		return
	}
	return
}
