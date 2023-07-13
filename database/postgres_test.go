package database

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bnjhope/pseudo-bank/transaction"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetTransactionsReturnsTransactionsOnSuccess(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected := make([]transaction.Transaction, 1)
	expected = append(expected, transaction.Transaction{
		Id:     1,
		Amount: 32.00,
		From:   "Ben",
		To:     "Emma",
	})

	result := sqlmock.NewRows([]string{"1", "32.00", "Ben", "Emma"})

	mock.ExpectQuery(regexp.QuoteMeta("select * from transaction")).WillReturnRows(result)

	pgDb := NewPgTransactionManager(db)

	actual, err := pgDb.GetTransactions()

	if err != nil {
		t.Error(err)
	}

	for ix, actual_row := range actual {
		expected_row := expected[ix]
		if expected_row != actual_row {
			t.Fatalf("Did not match rows\nExpected: %v\nActual: %v", expected_row, actual_row)
		}
	}
}

func TestGetTransactionsReturnsErrOnDbQueryFailure(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedMessage := "The query for transactions failed"

	mock.ExpectQuery(regexp.QuoteMeta("select * from transaction")).WillReturnError(fmt.Errorf(expectedMessage))

	pgDb := NewPgTransactionManager(db)

	actual, err := pgDb.GetTransactions()

	if actual != nil {
		t.Errorf("Recevied result back %v", actual)
	}

	if err == nil {
		t.Errorf("No err returned")
	}

	if err.Error() != expectedMessage {
		t.Errorf("No err returned")
	}
}
