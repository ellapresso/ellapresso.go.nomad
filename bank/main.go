package main

import (
	"fmt"

	accounts "github.com/ellapresso/bank/banking"
)

func main() {
	account := accounts.NewAccount("ellapresso")
	account2 := accounts.NewAccount("ellapresso2")
	account.Deposit(10)
	account2.Deposit(20)
	fmt.Println(account.Balance())
	fmt.Println(account2.Balance())

	// err := account.Withdraw(15)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance())

	fmt.Println(account.Balance(), account.Owner())
}
