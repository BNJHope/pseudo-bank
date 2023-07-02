package transaction

type Transaction struct {
	Id     int     `json:"id"`
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}
