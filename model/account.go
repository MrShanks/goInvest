package model

import (
	"github.com/MrShanks/goInvest/external"
	"github.com/google/uuid"
)

type Account struct {
	UUID     uuid.UUID
	Name     string
	Balance  float64
	Currency string
}

type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Accounts  []*Account
}

type Networth struct {
	Owner    *Person
	Balance  float64
	Currency string
}

func (n *Networth) CalculateBalance(accounts []*Account) float64 {
	var total float64
	for _, acc := range accounts {

		exchangeRate := 1.0
		if acc.Currency != n.Currency {
			chfRate := external.GetCurrentEchangeRate(n.Currency)
			otherRate := external.GetCurrentEchangeRate(acc.Currency)

			exchangeRate = chfRate * otherRate
		}

		total += acc.Balance * exchangeRate
	}

	return total
}
