package schema

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserOutput struct {
	ID       int     `json:"id"`
	UserName string  `json:"username"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
