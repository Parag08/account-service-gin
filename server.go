package main

import (
	"github.com/account-service-gin/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/status", api.Status)
	router.Run()
}
