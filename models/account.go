package models

import "time"

// Account : account struct of a customer
type Account struct {
	UserID    string     `json:"user_id"`
	Balance   uint       `json:"balance"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type CreateAccountParams struct {
	UserID string `json:"user_id"`
}
