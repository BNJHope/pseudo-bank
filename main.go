package main

import (
	"fmt"
	"io"
	"net/http"

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
	fmt.Printf("got /hello request\n")
	w.Header().Set("Content-Type", "application/json")

	transactions, nil := transaction.GetTransactions()
	io.WriteString(w, "Hello, HTTP!\n")
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
