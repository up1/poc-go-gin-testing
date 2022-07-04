package q1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Address_id string `json:"address_id"`
}

type Response struct {
	Message string `json:"message"`
}

func getFakeData() error {
	return nil
}

// ClaimRewards: Working with Post form
func ClaimRewards(c *gin.Context) {
	var address_id = c.PostForm("address_id")
	error := getFakeData() // Fake data

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": address_id,
	})
}

// ClaimRewards: Working with Post with JSON body
func ClaimRewardsJson(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var address_id = request.Address_id
	error := getFakeData() // Fake data

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": address_id,
	})
}
