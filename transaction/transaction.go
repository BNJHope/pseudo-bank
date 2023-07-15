package transaction

type Transaction struct {
	Id     int     `json:"id"`
	Amount float64 `json:"amount"`
	From   []byte  `json:"from"`
	To     []byte  `json:"to"`
}
