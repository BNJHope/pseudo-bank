//go:build integration
// +build integration

package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/bnjhope/pseudo-bank/transaction"
	"github.com/bnjhope/pseudo-bank/user"
	"github.com/google/go-cmp/cmp"
)

func TestIntegrationGetTransactionsReturnsTransactionsOnSuccess(t *testing.T) {

	expected := make([]transaction.Transaction, 0)
	expected = append(expected, transaction.Transaction{
		Id:     1,
		Amount: 20.5,
		From:   "d2e19190-59c8-4a43-8bb7-a729ea2b5173",
		To:     "1a8580b6-fb6c-4f3a-8254-3c19e638f385",
	})

	app_url := os.Getenv("APP_URL")
	transaction_url := fmt.Sprintf("http://%s/transaction", app_url)

	resp, httpGetErr := http.Get(transaction_url)

	if httpGetErr != nil {
		t.Errorf("HTTP GET /transaction err in test: %v", httpGetErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("HTTP GET /transaction resulted in >399 response: %v %v", resp.StatusCode, string(respBody))
	}

	var actual []transaction.Transaction

	jsonDecodeErr := json.NewDecoder(resp.Body).Decode(&actual)

	if jsonDecodeErr != nil {
		t.Fatalf("JSON Decode error in getting /transaction:\nResponse: %v\n %v", resp.Body, jsonDecodeErr)
	}

	if len(expected) != len(actual) {
		t.Fatalf("Number of transactions has different length.\n Expected Length %v\nActual Length: %v", len(expected), len(actual))
	}

	for ix, actual_row := range actual {
		expected_row := expected[ix]
		if !cmp.Equal(expected_row, actual_row) {
			t.Fatalf("Did not match rows\nExpected: %v\nActual: %v\nDiff: %v", expected_row, actual_row, cmp.Diff(expected_row, actual_row))
		}
	}
}

func TestIntegrationGetUserReturnsUserOnSuccess(t *testing.T) {
	var (
		userUrl    *url.URL
		urlErr     error
		testUserId string = "1a8580b6-fb6c-4f3a-8254-3c19e638f385"
	)

	expected := user.User{
		Id:        testUserId,
		FirstName: "Second",
		Surname:   "User",
		Balance:   0,
	}

	appUrl := os.Getenv("APP_URL")
	if userUrl, urlErr = url.Parse(fmt.Sprintf("http://%s/user", appUrl)); urlErr != nil {
		t.Fatalf("Error constructing get user URL: %v", urlErr)
	}

	userUrlVals := userUrl.Query()
	userUrlVals.Add("id", testUserId)
	userUrl.RawQuery = userUrlVals.Encode()

	resp, httpGetErr := http.Get(userUrl.String())

	if httpGetErr != nil {
		t.Errorf("HTTP GET /user err in test: %v", httpGetErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("HTTP GET /user resulted in >399 response: %v %v", resp.StatusCode, string(respBody))
	}

	var actual user.User

	jsonDecodeErr := json.NewDecoder(resp.Body).Decode(&actual)

	if jsonDecodeErr != nil {
		t.Fatalf("JSON Decode error in getting /transaction:\nResponse: %v\n %v", resp.Body, jsonDecodeErr)
	}

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Did not match rows\nExpected: %v\nActual: %v\nDiff: %v", expected, actual, cmp.Diff(expected, actual))
	}
}
