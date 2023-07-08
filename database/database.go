package database

import (
	"github.com/bnjhope/pseudo-bank/transaction"
)

type TransactionManager interface {
	GetTransactions() ([]transaction.Transaction, error)
	SaveTransaction(*transaction.Transaction) (int64, error)
}
