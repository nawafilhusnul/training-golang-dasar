package main

import (
	"errors"
	"fmt"
)

/*
===============================================================================
TASK 1: BANK ACCOUNT SYSTEM
===============================================================================
Instruksi:
1. Buat struct BankAccount dengan field: AccountNumber, Balance, Owner
2. Implementasikan method Deposit untuk menambah saldo
3. Implementasikan method Withdraw untuk mengambil uang (cek saldo cukup)
4. Implementasikan method GetBalance untuk melihat saldo
5. Implementasikan method Transfer untuk transfer ke account lain
6. Validasi: amount harus > 0, saldo cukup untuk withdraw

STARTER CODE:
*/

type BankAccount struct {
	// TODO: Definisikan field-field yang diperlukan
	AccountNumber string
	Balance       float64
	Owner         string
}

// TODO: Implementasikan method-method berikut:

// NewBankAccount adalah constructor untuk BankAccount
func NewBankAccount(accountNumber, owner string, initialBalance float64) (*BankAccount, error) {
	// IMPLEMENTASI DI SINI
	if accountNumber == "" || owner == "" {
		return nil, errors.New("account number dan owner tidak boleh kosong")
	}
	if initialBalance < 0 {
		return nil, errors.New("saldo awal tidak boleh negatif")
	}

	return &BankAccount{
		AccountNumber: accountNumber,
		Balance:       initialBalance,
		Owner:         owner,
	}, nil
}

// Deposit menambahkan uang ke account
func (ba *BankAccount) Deposit(amount float64) error {
	// IMPLEMENTASI DI SINI
	if amount <= 0 {
		return errors.New("amount harus lebih besar dari 0")
	}
	ba.Balance += amount
	return nil
}

// Withdraw mengambil uang dari account
func (ba *BankAccount) Withdraw(amount float64) error {
	// IMPLEMENTASI DI SINI
	if amount <= 0 {
		return errors.New("amount harus lebih besar dari 0")
	}
	if ba.Balance < amount {
		return errors.New("saldo tidak cukup")
	}
	ba.Balance -= amount
	return nil
}

// GetBalance mengembalikan saldo saat ini
func (ba BankAccount) GetBalance() float64 {
	// IMPLEMENTASI DI SINI
	return ba.Balance
}

// Transfer mentransfer uang ke account lain
func (ba *BankAccount) Transfer(to *BankAccount, amount float64) error {
	// IMPLEMENTASI DI SINI
	if to == nil {
		return errors.New("target account tidak boleh nil")
	}
	if amount <= 0 {
		return errors.New("amount harus lebih besar dari 0")
	}
	if ba.Balance < amount {
		return errors.New("saldo tidak cukup")
	}

	ba.Balance -= amount
	to.Balance += amount
	return nil
}

// GetAccountInfo mengembalikan informasi account
func (ba BankAccount) GetAccountInfo() string {
	// IMPLEMENTASI DI SINI
	return fmt.Sprintf("Account: %s, Owner: %s, Balance: Rp %.2f",
		ba.AccountNumber, ba.Owner, ba.Balance)
}

func testBankAccount() {
	fmt.Println("TASK: Bank Account System")
	acc1, _ := NewBankAccount("001", "Alice", 1000000)
	acc2, _ := NewBankAccount("002", "Bob", 500000)

	acc1.Deposit(200000)
	acc1.Withdraw(150000)
	acc1.Transfer(acc2, 100000)

	fmt.Println(acc1.GetAccountInfo())
	fmt.Println(acc2.GetAccountInfo())
	fmt.Println()

}
