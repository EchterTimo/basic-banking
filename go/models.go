package main

import "github.com/google/uuid"

// region Account

type Account struct {
	id        string
	ownerName string
	balance   float64
}

func (a Account) GetID() string {
	return a.id
}

func (a Account) GetBalance() float64 {
	return a.balance
}

func (a Account) HasSufficientBalance(amount float64) bool {
	if a.balance >= amount {
		return true
	}
	return false
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	if !a.HasSufficientBalance(amount) {
		return false
	}
	a.balance -= amount
	return true
}

func (a *Account) Transfer(amount float64, target *Account) bool {
	if !a.HasSufficientBalance(amount) {
		return false
	}
	a.balance -= amount
	target.balance += amount
	return true
}

// endregion

// region Bank

type Bank struct {
	bankName string
	accounts []Account
}

func NewBank(bankName string) Bank {
	return Bank{
		bankName: bankName,
		accounts: []Account{},
	}
}

func (b *Bank) CreateAccount(ownerName string) Account {
	// todo: auto generate custom id
	newAccount := Account{id: uuid.New().String(), ownerName: ownerName, balance: 0.0}
	b.accounts = append(b.accounts, newAccount)
	return newAccount
}

func (b *Bank) GetAccountById(id string) *Account {
	for i := range b.accounts {
		if b.accounts[i].id == id {
			return &b.accounts[i]
		}
	}
	return nil
}

// endregion
