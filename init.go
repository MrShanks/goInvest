package main

import "github.com/MrShanks/goInvest/model"

func InitNetworth() *model.Networth {
	simone := &model.Person{
		Firstname: "Simone",
		Lastname:  "Staffoli",
		Email:     "simonestaffoli@gmail.com",
	}

	degiro := model.Account{
		Name:     "Degiro",
		Balance:  92130.17,
		Currency: "CHF",
	}

	ubsMain := model.Account{
		Name:     "UBS Main",
		Balance:  21889.92,
		Currency: "CHF",
	}

	ubsCommon := model.Account{
		Name:     "UBS Common",
		Balance:  203.17,
		Currency: "CHF",
	}

	ubsSavings := model.Account{
		Name:     "UBS Savings",
		Balance:  15279.7,
		Currency: "CHF",
	}

	viac := model.Account{
		Name:     "Viac 3a Pillar",
		Balance:  30555.67,
		Currency: "CHF",
	}

	bnl := model.Account{
		Name:     "BNL",
		Balance:  297.78,
		Currency: "EUR",
	}

	simone.Accounts = append(simone.Accounts, &degiro)
	simone.Accounts = append(simone.Accounts, &ubsMain)
	simone.Accounts = append(simone.Accounts, &ubsCommon)
	simone.Accounts = append(simone.Accounts, &ubsSavings)
	simone.Accounts = append(simone.Accounts, &viac)
	simone.Accounts = append(simone.Accounts, &bnl)

	nw := model.Networth{
		Owner:    simone,
		Currency: "CHF",
	}

	nw.Balance = nw.CalculateBalance(nw.Owner.Accounts)

	return &nw
}
