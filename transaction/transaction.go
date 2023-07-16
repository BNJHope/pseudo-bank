package transaction

import "fmt"

type Transaction struct {
	Id     int     `json:"id"`
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type NotEnoughFundsInAccountError struct {
	AccountBalance    float64
	TransactionAmount float64
}

func (err *NotEnoughFundsInAccountError) Error() string {
	return fmt.Sprintf(`Unable to process transaction, not enough funds in account
	Account Balance: %f,
	Transaction Amount: %f`, err.AccountBalance, err.TransactionAmount)
}
