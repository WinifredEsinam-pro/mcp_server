package main

import "fmt"

type Account struct {
	AccountNumber string
	Balance       float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Println("Deposited:", amount)
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		fmt.Println("Insufficient balance")
		return
	}
	a.Balance -= amount
	fmt.Println("Withdrawn:", amount)
}

func (a Account) Display() {
	fmt.Println("Account:", a.AccountNumber)
	fmt.Println("Balance:", a.Balance)
	fmt.Println("----------------------")
}

func main() {
	acc1 := Account{"WIUC-0001", 5000}
	acc2 := Account{"WIUC-0002", 2000}

	fmt.Println("Initial Accounts:")
	acc1.Display()
	acc2.Display()

	acc1.Deposit(1000)
	acc2.Withdraw(500)

	fmt.Println("Updated Accounts:")
	acc1.Display()
	acc2.Display()
}