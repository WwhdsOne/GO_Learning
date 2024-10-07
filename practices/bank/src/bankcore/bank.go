package bankcore

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return errors.New("the amount is bigger than balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(a1, a2 *Account, money float64) error {
	if money <= 0 {
		return errors.New("the money amount should be greater than zero")
	}
	if money > a.Balance {
		return errors.New("转账账户余额不足")
	}
	a1.Balance -= money
	a2.Balance += money
	return nil
}
