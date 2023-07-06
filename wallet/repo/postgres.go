package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"julo-mini-wallet/wallet"
	"strconv"

	sq "github.com/Masterminds/squirrel"
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

func (p *repositoryStruct) GetAccount(userID string) (walletData wallet.Wallet, err error) {
	placeholder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	psql := placeholder.Select("id", "owned_by", "status", "balance", "enabled_at").From("wallet")
	whereSQL := psql.Where(sq.Eq{"owned_by": userID})

	rows, err := whereSQL.RunWith(p.pqConnections).Query()
	if err != nil {
		return walletData, err
	}
	if rows != nil {
		defer rows.Close()
		if !rows.Next() {
			return walletData, wallet.ErrAccountNotExist
		}
		for rows.Next() {
			var rowID int
			err := rows.Scan(&rowID, &walletData.OwnedBy, &walletData.Status, &walletData.Balance, &walletData.EnabledAt)
			if err != nil {
				return walletData, err
			}
			walletData.ID = hashAndFormatID(rowID)
		}
	}
	return walletData, nil
}

func (p *repositoryStruct) CreateAccount(userID string) (err error) {

	stmt, err := p.pqConnections.Prepare("INSERT INTO wallet (owned_by,status,balance)  VALUES ($1, $2, $3)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(userID, "disabled", 0)
	if err != nil {
		return
	}

	return
}

func hashAndFormatID(id int) string {

	idBytes := []byte(strconv.Itoa(id))

	hash := sha1.Sum(idBytes)

	hashString := hex.EncodeToString(hash[:])

	formattedID := fmt.Sprintf("%s-%s-%s-%s-%s",
		hashString[:8], hashString[8:12], hashString[12:16], hashString[16:20], hashString[20:])

	return formattedID
}
