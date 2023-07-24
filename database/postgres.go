package database

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/bnjhope/pseudo-bank/transaction"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PgTransactionsManager struct {
	db *sql.DB
}

func NewPgTransactionManager(db *sql.DB) PgTransactionsManager {
	return PgTransactionsManager{db: db}
}

func (tm PgTransactionsManager) GetTransactions(userId string) ([]transaction.Transaction, error) {
	var (
		rows       *sql.Rows
		dbQueryErr error
	)

	if userId != "" {
		rows, dbQueryErr = tm.db.Query("select * from transactions where transactions.fromaccount = $1", userId)
	} else {
		rows, dbQueryErr = tm.db.Query("select * from transactions")
	}

	if dbQueryErr != nil {
		return nil, dbQueryErr
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

func (tm PgTransactionsManager) SaveTransaction(t *transaction.Transaction) (int64, error) {
	var (
		Id                    int64
		NewAccountFromBalance float64
	)

	tx, beginTxErr := tm.db.Begin()

	if beginTxErr != nil {
		log.Fatal().Err(beginTxErr).Msg("Failed to begin transaction")
		return -1, beginTxErr
	}

	insertTransactionErr := tx.QueryRow(`INSERT into transactions (Amount, FromAccount, ToAccount)
	VALUES ($1, $2, $3)
	RETURNING id`, t.Amount, t.From, t.To).Scan(&Id)

	if insertTransactionErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(rollbackErr).Msg("Failed to rollback insert transaction")
		}
		return -1, fmt.Errorf("insert transaction: %w", insertTransactionErr)
	}

	decrementBalanceErr := tx.QueryRow(`UPDATE users
	SET balance = users.balance - $1
	WHERE users.id = $2
	RETURNING users.balance`, t.Amount, t.From).Scan(&NewAccountFromBalance)

	if decrementBalanceErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(rollbackErr).Msg("Failed to rollback decrement balance transaction")
		}
		return -1, fmt.Errorf("decrement balance: %w", decrementBalanceErr)
	}

	if NewAccountFromBalance < 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(rollbackErr).Msg("Failed to rollback from not enough funds in account")
		}
		return -1, &transaction.NotEnoughFundsInAccountError{
			AccountBalance:    NewAccountFromBalance + t.Amount,
			TransactionAmount: t.Amount,
		}
	}

	incrementBalanceErr := tx.QueryRow(`UPDATE users
	SET balance = users.balance + $1
	WHERE users.id = $2
	RETURNING users.balance`, t.Amount, t.To).Scan(&NewAccountFromBalance)

	if incrementBalanceErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal().Err(rollbackErr).Msg("Failed to rollback increment balance transaction")
		}
		return -1, fmt.Errorf("increment balance: %w", incrementBalanceErr)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		log.Fatal().Err(commitErr).Msg("Failed to commit transaction")
		return -1, commitErr
	}

	return Id, nil
}
