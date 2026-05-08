package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(ptr1 *customer, ptr2 transaction) error {

	if ptr1.id != ptr2.customerID {
		return errors.New("customer ID mismatch")
	}

	if ptr2.transactionType == transactionDeposit {
		ptr1.balance += ptr2.amount
		return nil
	}

	if ptr2.transactionType == transactionWithdrawal {

		if ptr1.balance < ptr2.amount {
			return errors.New("insufficient funds")
		}

		ptr1.balance -= ptr2.amount
		return nil
	}

	return errors.New("unknown transaction type")
}
