package store

import (
	"database/sql"
	"fmt"

	"github.com/MrShanks/goInvest/model"
)

type AccountStore struct {
	db *sql.DB
}

func (a *AccountStore) New(acc *model.Account) {
	q := fmt.Sprintf("INSERT INTO account (name, balance, currency) values (%s,%f,%s);", acc.Name, acc.Balance, acc.Currency)

	_, err := a.db.Query(q)
	if err != nil {
		fmt.Printf("Couldn't insert new account: %v\n", err)
	}
}
