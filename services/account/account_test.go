package account_test

import (
	"testing"

	"github.com/account-service-gin/models"
	"github.com/account-service-gin/services/account"
)

func initDB() {
	// Setup db
}

func resetDB() {
	// Reset db
}

func TestCreateAccount(t *testing.T) {
	var (
		err error
		acc *models.Account
	)
	if acc, err = account.CreateAccount(models.CreateAccountParams{
		UserID: "a7f39653-025d-4932-bb3a-93c5c256587a",
	}); err != nil {
		t.Fail()
	} else {
		t.Log("Account creation sucessful:", *acc)
	}

}
