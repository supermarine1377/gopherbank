package domain

type Transaction struct {
	ID     int `json:"id,omitempty"`
	Amount int `json:"amount,omitempty"`
	UserID int `json:"user_id,omitempty"`
}
