package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type (
	DatabaseObjects struct {
		Master *sqlx.DB
	}
)

// GetDatabaseConns  ...
func GetDatabaseConns() *DatabaseObjects {

	dbMaster, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "randy", "password", "fullstack-postgres:5432", "julo_db"))
	if err != nil {
		log.Fatal(err, " failed connect db")
	}
	dbMaster.SetMaxOpenConns(30)
	dbMaster.SetMaxIdleConns(5)

	dbConns := &DatabaseObjects{
		Master: dbMaster,
	}

	return dbConns
}
