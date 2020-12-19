package account

import (
	"github.com/account-service-gin/database"
	"github.com/account-service-gin/models"
)

func CreateAccount(createAccountParams models.CreateAccountParams) (models.Account, error) {
	var (
		account models.Account
		err     error
	)
	db := database.NewStore()
	rows, err := db.Query(`INSERT INTO accounts (user_id)
			  VALUES ($1)
			  RETURNING user_id, balance, created_at, updated_at
	`, createAccountParams.UserID)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&account.UserID, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return account, err
		}
	}
	if err = rows.Err(); err != nil {
		return account, err
	}
	return account, err
}
