package database

import (
	"database/sql"
	"os"

	"github.com/bnjhope/pseudo-bank/transaction"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type TransactionManager interface {
	getTransactions() []transaction.Transaction
}

func GetTransactions() ([]transaction.Transaction, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, dbQueryErr := db.Query("select * from transaction")
	if dbQueryErr != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := make([]transaction.Transaction, 0)
	for rows.Next() {
		var (
			id                     int
			Amount                 float64
			FromAccount, ToAccount string
		)
		if err := rows.Scan(&id, &Amount, &FromAccount, &ToAccount); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction.Transaction{Id: id, Amount: Amount, From: FromAccount, To: ToAccount})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
