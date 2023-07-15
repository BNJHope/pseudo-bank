package database

import (
	"database/sql"

	"github.com/bnjhope/pseudo-bank/transaction"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PgTransactionsManager struct {
	db *sql.DB
}

func NewPgTransactionManager(db *sql.DB) PgTransactionsManager {
	return PgTransactionsManager{db: db}
}

func (tm PgTransactionsManager) GetTransactions() ([]transaction.Transaction, error) {
	rows, dbQueryErr := tm.db.Query("select * from transactions")
	if dbQueryErr != nil {
		return nil, dbQueryErr
	}

	defer rows.Close()

	transactions := make([]transaction.Transaction, 0)
	for rows.Next() {
		var (
			id                     int
			Amount                 float64
			FromAccount, ToAccount []byte
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

func (tm PgTransactionsManager) SaveTransaction(t *transaction.Transaction) (int64, error) {
	var Id int64
	queryDbErr := tm.db.QueryRow(`INSERT into transactions (Amount, FromAccount, ToAccount)
	VALUES ($1, $2, $3)
	RETURNING id`, t.Amount, t.From, t.To).Scan(&Id)

	if queryDbErr != nil {
		return -1, queryDbErr
	}

	return Id, nil
}
