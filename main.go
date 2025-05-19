package main

import (
	"fmt"

	"github.com/MrShanks/goInvest/model"
	"github.com/MrShanks/goInvest/storage"
	"github.com/MrShanks/goInvest/store"
	_ "github.com/lib/pq"
)

func main() {

	db := storage.Connect()
	defer db.Close()

	accStore := store.AccountStore{
		DB: db,
	}

	nw := model.Networth{
		Owner:    "simonestaffoli@gmail.com",
		Currency: "CHF",
		Accounts: accStore.Get(),
	}

	nw.CalculateBalance()

	fmt.Printf("Welcome to your dashboard %s\n", nw.Owner)
	fmt.Printf("Your total Networth is: %.2f CHF\n", nw.Balance)
}
