package wallet

import (
	"errors"
	"time"
)

const (
	StatusDisabled = "disabled"
	StatusEnabled  = "enabled"

	TypeWithdrawal = "withdrawal"
	TypeDeposit    = "deposit"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrAccountNotExist     = errors.New("account not exist")
)

type (
	NewAccountResponse struct {
		Token string `json:"token"`
	}
	EnableWalletResponse struct {
		ID        string     `json:"id"`
		OwnedBy   string     `json:"owned_by"`
		Status    string     `json:"status"`
		EnabledAt *time.Time `json:"enabled_at"`
		Balance   float64    `json:"balance"`
	}
	DisableWalletResponse struct {
		ID         string     `json:"id"`
		OwnedBy    string     `json:"owned_by"`
		Status     string     `json:"status"`
		DisabledAt *time.Time `json:"disabled_at"`
		Balance    float64    `json:"balance"`
	}
	Wallet struct {
		ID         string     `json:"id"`
		OwnedBy    string     `json:"owned_by"`
		Status     string     `json:"status"`
		EnabledAt  *time.Time `json:"enabled_at"`
		DisabledAt *time.Time `json:"disabled_at"`
		Balance    float64    `json:"balance"`
	}
	Transaction struct {
		ID               string    `json:"id"`
		WalletID         string    `json:"wallet_id"`
		ActorID          string    `json:"actor_id"`
		Status           string    `json:"status"`
		TransactedAt     time.Time `json:"transacted_at"`
		Type             string    `json:"type"`
		Amount           float64   `json:"amount"`
		ReferenceID      string    `json:"reference_id"`
		InputReferenceID string    `json:"input_reference_id"`
	}

	TransactionResponse struct {
		ID           string    `json:"id"`
		Status       string    `json:"status"`
		TransactedAt time.Time `json:"transacted_at"`
		Type         string    `json:"type"`
		Amount       float64   `json:"amount"`
		ReferenceID  string    `json:"reference_id"`
	}

	Deposit struct {
		ID          string    `json:"id"`
		WalletID    string    `json:"wallet_id"`
		DepositedBy string    `json:"deposited_by"`
		Status      string    `json:"status"`
		DepositedAt time.Time `json:"deposited_at"`
		Amount      float64   `json:"amount"`
		ReferenceID string    `json:"reference_id"`
	}
	WithDrawal struct {
		ID          string    `json:"id"`
		WalletID    string    `json:"wallet_id"`
		WithdrawnBy string    `json:"withdrawn_by"`
		Status      string    `json:"status"`
		WithdrawnAt time.Time `json:"withdrawn_at"`
		Amount      float64   `json:"amount"`
		ReferenceID string    `json:"reference_id"`
	}
)

type (
	Repository interface {
		CreateAccount(userID string) (err error)
		GetAccount(userID string) (walletData Wallet, err error)
		UpdateAccountStatus(userID string, status string) (walletData Wallet, err error)
		InsertTransactionData(transactionInput Transaction) (transactionID string, err error)
		GetTransactionData(walletID string) (transactionsData []TransactionResponse, err error)
		UpdateAccountBalance(userID string, balance float64) (err error)
	}
	Usecase interface {
		CreateAccount(userID string) (userToken NewAccountResponse, err error)
		EnableWallet(userID string) (walletData EnableWalletResponse, err error)
		DisableWallet(userID string) (walletData DisableWalletResponse, err error)
		GetWalletBalance(userID string) (walletData EnableWalletResponse, err error)
		Deposit(userID string, amount float64, referenceID string) (depositData Deposit, err error)
		Withdraw(userID string, amount float64, referenceID string) (withdrawData WithDrawal, err error)
		GetWalletTransactions(userID string) (transactionsData []TransactionResponse, err error)
	}
)
