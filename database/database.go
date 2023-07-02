package database

import (
	"database/sql"
	"fmt"
	"log"
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	rows, dbQueryErr := db.Query("select * from transaction")
	if dbQueryErr != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
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
			log.Fatal(err)
		}
		transactions = append(transactions, transaction.Transaction{Id: id, Amount: Amount, From: FromAccount, To: ToAccount})
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return transactions, nil
}
