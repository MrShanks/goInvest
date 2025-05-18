package main

import "github.com/MrShanks/goInvest/model"

func InitNetworth() *Networth {
	simone := &model.Person{
		Firstname: "Simone",
		Lastname:  "Staffoli",
		Email:     "simonestaffoli@gmail.com",
	}

	degiro := model.Account{
		Name:     "Degiro",
		Total:    92130.17,
		Currency: "CHF",
	}

	ubsMain := model.Account{
		Name:     "UBS Main",
		Total:    21889.92,
		Currency: "CHF",
	}

	ubsCommon := model.Account{
		Name:     "UBS Common",
		Total:    203.17,
		Currency: "CHF",
	}

	ubsSavings := model.Account{
		Name:     "UBS Savings",
		Total:    15279.7,
		Currency: "CHF",
	}

	viac := model.Account{
		Name:     "Viac 3a Pillar",
		Total:    30555.67,
		Currency: "CHF",
	}

	bnl := model.Account{
		Name:     "BNL",
		Total:    297.78,
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

	nw.Total = nw.CalculateTotal(nw.Owner.Accounts)

	return &nw
}
