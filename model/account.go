package model

import (
	"github.com/MrShanks/goInvest/exchange"
	"github.com/google/uuid"
)

type Account struct {
	UUID     uuid.UUID
	Name     string
	Balance  float64
	Currency string
}

type Networth struct {
	Owner    string
	Balance  float64
	Currency string
	Accounts []*Account
}

func (n *Networth) CalculateBalance() {
	var total float64
	for _, acc := range n.Accounts {

		exchangeRate := 1.0
		if acc.Currency != n.Currency {
			chfRate := exchange.GetCurrentEchangeRate(n.Currency)
			otherRate := exchange.GetCurrentEchangeRate(acc.Currency)

			exchangeRate = chfRate * otherRate
		}

		total += acc.Balance * exchangeRate
	}

	n.Balance = total
}
