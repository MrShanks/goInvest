package store

import (
	"database/sql"
	"fmt"

	"github.com/MrShanks/goInvest/model"
)

type AccountStore struct {
	DB *sql.DB
}

func (a *AccountStore) New(acc *model.Account) {
	q := fmt.Sprintf("INSERT INTO account (name, balance, currency) values (%s,%f,%s);", acc.Name, acc.Balance, acc.Currency)

	_, err := a.DB.Query(q)
	if err != nil {
		fmt.Printf("Couldn't insert new account: %v\n", err)
	}
}

func (a *AccountStore) Get() []*model.Account {
	accs := []*model.Account{}

	q := "SELECT * FROM account;"

	rows, err := a.DB.Query(q)
	if err != nil {
		fmt.Printf("Couldn't insert new account: %v\n", err)
	}

	for rows.Next() {
		var acc model.Account
		if err := rows.Scan(&acc.UUID, &acc.Name, &acc.Balance, &acc.Currency); err != nil {
			fmt.Printf("Couldn't scan row: %v\n", err)
		}

		accs = append(accs, &acc)
	}

	return accs
}
