package domain

type User struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Balance   int    `json:"balance,omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty"`
}
