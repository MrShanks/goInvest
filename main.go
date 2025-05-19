package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/MrShanks/goInvest/model"
	"github.com/MrShanks/goInvest/storage"
	"github.com/MrShanks/goInvest/store"
	_ "github.com/lib/pq"
)

func main() {
	simone := &model.Person{
		Firstname: "Simone",
		Lastname:  "Staffoli",
		Email:     "simonestaffoli@gmail.com",
	}

	nw := model.Networth{
		Owner:    simone,
		Currency: "CHF",
	}

	fmt.Printf("Welcome to your dashboard %s %s\n", nw.Owner.Firstname, nw.Owner.Lastname)
	fmt.Printf("Your total Networth is: %.2f CHF\n", nw.Balance)

	db := storage.Connect()
	defer db.Close()

	accStore := store.AccountStore{
		DB: db,
	}

	accounts := accStore.Get()

	for _, acc := range accounts {
		fmt.Println(*acc)
	}

	nw.Owner.Accounts = accounts

	nw.CalculateBalance()

	fmt.Println(nw.Balance)
}

func Update(nw *model.Networth) {
	input := bufio.NewScanner(os.Stdin)

	for _, acc := range nw.Owner.Accounts {
		fmt.Printf("Account: %s\n", acc.Name)

		fmt.Printf("Insert new account balance: ")
		input.Scan()
		newBalance := input.Text()
		fmt.Printf("\n")

		nb, err := strconv.ParseFloat(newBalance, 64)

		acc.Balance = nb
		if err != nil {
			fmt.Printf("please insert a valid float number\n")
			continue
		}
	}

	nw.CalculateBalance()
	fmt.Printf("New Balance: %.2f\n", nw.Balance)
}
