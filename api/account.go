package api

import (
	"net/http"

	"github.com/account-service-gin/models"
	"github.com/account-service-gin/services/account"
	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(ctx *gin.Context) {
	var (
		err  error
		resp struct {
			Account models.Account `json:"account"`
		}
		createAccountParams models.CreateAccountParams
	)
	if err := ctx.ShouldBindJSON(&createAccountParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp.Account, err = account.CreateAccount(createAccountParams); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"account": resp.Account})
}
