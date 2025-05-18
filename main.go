package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MrShanks/goInvest/model"
	"github.com/MrShanks/goInvest/storage"
	_ "github.com/lib/pq"
)

func main() {

	nw := InitNetworth()

	fmt.Printf("Welcome to your dashboard %s %s\n", nw.Owner.Firstname, nw.Owner.Lastname)
	fmt.Printf("Your total Networth is: %.2f CHF\n", nw.Balance)

	db := storage.Connect()
	defer db.Close()

	// Query the account table
	rows, err := db.Query("SELECT name, balance, currency FROM account")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	rows.Next()

	acc := &model.Account{}
	if err := rows.Scan(&acc.Name, &acc.Balance, &acc.Currency); err != nil {
		fmt.Println(err)
	}

	fmt.Println(*acc)
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

	nw.CalculateBalance(nw.Owner.Accounts)
	fmt.Printf("New Balance: %.2f\n", nw.Balance)
}
