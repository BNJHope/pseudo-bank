package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/bnjhope/pseudo-bank/transaction"
)

type TransactionManager interface {
	getTransactions() []transaction.Transaction
}

func GetTransactions() ([]transaction.Transaction, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	transactions := make([]transaction.Transaction, 10)
	err = db.QueryRowContext(context.Background(), "select * from transactions where id=$1").Scan(&transactions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return transactions, nil
}
