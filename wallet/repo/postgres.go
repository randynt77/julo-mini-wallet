package repository

import (
	"fmt"
	"julo-mini-wallet/wallet"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	repositoryStruct struct {
		pqConnections *sqlx.DB
	}
)

func NewRepo(master *sqlx.DB) wallet.Repository {
	return &repositoryStruct{
		pqConnections: master,
	}
}

func (p *repositoryStruct) UpdateAccountBalance(userID string, balance float64) (err error) {

	stmt, err := p.pqConnections.Prepare("UPDATE wallet set balance = $1 where owned_by = $2")
	if err != nil {
		return
	}

	_, err = stmt.Exec(balance, userID)
	if err != nil {
		return
	}

	return nil
}

func (p *repositoryStruct) UpdateAccountStatus(userID string, status string) (walletData wallet.Wallet, err error) {
	field := "enabled_at"
	if status == wallet.StatusDisabled {
		field = "disabled_at"
	}

	query := fmt.Sprintf("UPDATE wallet set status = $1, %s = $2 where owned_by = $3 Returning id,owned_by,status,enabled_at,disabled_at,balance", field)
	stmt, err := p.pqConnections.Prepare(query)
	if err != nil {
		return
	}

	err = stmt.QueryRow(status, time.Now(), userID).Scan(&walletData.ID, &walletData.OwnedBy, &walletData.Status, &walletData.EnabledAt, &walletData.DisabledAt, &walletData.Balance)
	if err != nil {
		return
	}

	return walletData, nil
}

func (p *repositoryStruct) GetAccount(userID string) (walletData wallet.Wallet, err error) {

	stmt, err := p.pqConnections.Prepare("SELECT id,owned_by,status,balance,enabled_at,disabled_at from wallet where owned_by = $1")
	if err != nil {
		return
	}

	err = stmt.QueryRow(userID).
		Scan(&walletData.ID, &walletData.OwnedBy, &walletData.Status, &walletData.Balance, &walletData.EnabledAt, &walletData.DisabledAt)

	return walletData, nil
}

func (p *repositoryStruct) CreateAccount(userID string) (err error) {

	stmt, err := p.pqConnections.Prepare("INSERT INTO wallet (id,owned_by,status,balance)  VALUES ($1, $2, $3, $4)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(uuid.New().String(), userID, "disabled", 0)
	if err != nil {
		return
	}
	return
}

func (p *repositoryStruct) InsertTransactionData(transactionInput wallet.Transaction) (transactionID string, err error) {

	stmt, err := p.pqConnections.Prepare("INSERT INTO transaction (id,wallet_id,actor_id,status,transacted_at,type,amount,reference_id,input_reference_id)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id")
	if err != nil {
		return
	}
	err = stmt.QueryRow(uuid.New().String(), transactionInput.WalletID, transactionInput.ActorID, transactionInput.Status, transactionInput.TransactedAt, transactionInput.Type, transactionInput.Amount, transactionInput.ReferenceID, transactionInput.InputReferenceID).Scan(&transactionID)
	if err != nil {
		return
	}

	return
}

func (p *repositoryStruct) GetTransactionData(walletID string) (transactionsData []wallet.TransactionResponse, err error) {

	stmt, err := p.pqConnections.Prepare("SELECT id,status,transacted_at,type,amount,reference_id from transaction where wallet_id = $1")
	if err != nil {
		return
	}
	rows, err := stmt.Query(walletID)
	if err != nil {
		return
	}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var transactionData wallet.TransactionResponse

			err := rows.Scan(&transactionData.ID, &transactionData.Status, &transactionData.TransactedAt, &transactionData.Type, &transactionData.Amount, &transactionData.ReferenceID)
			if err != nil {
				return transactionsData, err
			}

			transactionsData = append(transactionsData, transactionData)
		}
	}

	return transactionsData, nil
}
