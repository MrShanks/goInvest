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
	a.db.Query(q)
}
