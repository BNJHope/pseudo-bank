//go:build integration
// +build integration

package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/bnjhope/pseudo-bank/transaction"
	"github.com/google/go-cmp/cmp"
)

func TestIntegrationGetTransactionsReturnsTransactionsOnSuccess(t *testing.T) {

	expected := make([]transaction.Transaction, 0)
	expected = append(expected, transaction.Transaction{
		Id:     1,
		Amount: 20.5,
		From:   []byte("d2e19190-59c8-4a43-8bb7-a729ea2b5173"),
		To:     []byte("1a8580b6-fb6c-4f3a-8254-3c19e638f385"),
	})

	app_url := os.Getenv("APP_URL")
	transaction_url := fmt.Sprintf("http://%s/transaction", app_url)

	resp, httpGetErr := http.Get(transaction_url)

	if httpGetErr != nil {
		t.Errorf("HTTP GET /transaction err in test: %v", httpGetErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("HTTP GET /transaction resulted in >399 response: %v %v", resp.StatusCode, string(respBody))
	}

	var actual []transaction.Transaction

	jsonDecodeErr := json.NewDecoder(resp.Body).Decode(&actual)

	if jsonDecodeErr != nil {
		t.Fatalf("JSON Decode error in getting /transaction:\nResponse: %v\n %v", resp.Body, jsonDecodeErr)
	}

	for ix, actual_row := range actual {
		expected_row := expected[ix]
		if !cmp.Equal(expected_row, actual_row) {
			t.Fatalf("Did not match rows\nExpected: %v\nActual: %v\nDiff: %v", expected_row, actual_row, cmp.Diff(expected_row, actual_row))
		}
	}
}
