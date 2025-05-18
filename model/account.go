package store

import "github.com/google/uuid"

type Account struct {
	UUID     uuid.UUID
	Name     string
	Total    float64
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
	Total    float64
	Currency string
}

func (n *Networth) CalculateTotal(accounts []*Account) float64 {
	var total float64
	for _, acc := range accounts {

		exchangeRate := 1.0
		if acc.Currency != n.Currency {
			chfRate := GetCurrentEchangeRate(n.Currency)
			otherRate := GetCurrentEchangeRate(acc.Currency)

			exchangeRate = chfRate * otherRate
		}

		total += acc.Total * exchangeRate
	}

	return total
}
