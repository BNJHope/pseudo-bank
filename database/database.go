package database

import (
	"github.com/bnjhope/pseudo-bank/transaction"
	"github.com/bnjhope/pseudo-bank/user"
)

type TransactionManager interface {
	GetTransactions(userId string) ([]transaction.Transaction, error)
	SaveTransaction(*transaction.Transaction) (int64, error)
	GetUser(userId string) (user.User, error)
}
