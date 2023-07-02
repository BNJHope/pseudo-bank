package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bnjhope/pseudo-bank/database"
	"github.com/bnjhope/pseudo-bank/transaction"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		transactions []transaction.Transaction
		err          error
	)
	fmt.Printf("got /hello request\n")
	w.Header().Set("Content-Type", "application/json")

	transactions, err = database.GetTransactions()

	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, "my own error message", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(transactions)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/transactions", getTransactions)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
