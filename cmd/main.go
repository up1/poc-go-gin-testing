package main

import (
	"q1"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/claim", q1.ClaimRewards) // Working with Post form
	r.POST("/claim/json", q1.ClaimRewardsJson) // Working with Post JSON body
	r.Run()
}