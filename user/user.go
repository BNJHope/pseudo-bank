package user

type User struct {
	Id        string  `json:"id"`
	FirstName string  `json:"firstname"`
	Surname   string  `json:"surname"`
	Balance   float64 `json:"balance"`
}
