package wallet

import (
	"errors"
	"time"
)

const (
	TopupSourceAccount int = -999
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrAccountNotExist     = errors.New("account not exist")
)

type (
	NewAccountResponse struct {
		Token string
	}

	Wallet struct {
		ID        string    `json:"id,omitempty"`
		OwnedBy   string    `json:"owned_by,omitempty"`
		Status    string    `json:"status,omitempty"`
		EnabledAt time.Time `json:"enabled_at,omitempty"`
		Balance   float64   `json:"balance,omitempty"`
	}
	Transaction struct {
		ID           string    `json:"id,omitempty"`
		Status       string    `json:"status,omitempty"`
		TransactedAt time.Time `json:"transacted_at,omitempty"`
		Type         string    `json:"type,omitempty"`
		Amount       float64   `json:"amount,omitempty"`
		ReferenceID  string    `json:"reference_id,omitempty"`
	}
	Deposit struct {
		ID          string    `json:"id,omitempty"`
		DepositedBy string    `json:"deposited_by,omitempty"`
		Status      string    `json:"status,omitempty"`
		DepositedAt time.Time `json:"deposited_at,omitempty"`
		Amount      float64   `json:"amount,omitempty"`
		ReferenceID string    `json:"reference_id,omitempty"`
	}
	WithDrawal struct {
		ID          string    `json:"id,omitempty"`
		WithdrawnBy string    `json:"withdrawn_by,omitempty"`
		Status      string    `json:"status,omitempty"`
		WithdrawnAt time.Time `json:"withdrawn_at,omitempty"`
		Amount      float64   `json:"amount,omitempty"`
		ReferenceID string    `json:"reference_id,omitempty"`
	}
)

type (
	Repository interface {
		CreateAccount(userID string) (err error)
		GetAccount(userID string) (walletData Wallet, err error)
	}
	Usecase interface {
		CreateAccount(userID string) (userToken NewAccountResponse, err error)
	}
)
