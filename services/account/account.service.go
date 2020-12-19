package account

import (
	"github.com/account-service-gin/database"
	"github.com/account-service-gin/models"
)

// CreateAccount : create account in database
func CreateAccount(createAccountParams models.CreateAccountParams) (*models.Account, error) {
	var (
		acc models.Account
		err error
	)
	db := database.NewStore()
	rows, err := db.Query(`INSERT INTO accounts (user_id)
			  VALUES ($1)
			  RETURNING user_id, balance, created_at, updated_at
	`, createAccountParams.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&acc.UserID, &acc.Balance, &acc.CreatedAt, &acc.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &acc, err
}
