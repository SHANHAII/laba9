package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskRequest struct {
	Data string `json:"data" binding:"required"`
}

func main() {
	r := gin.Default()

)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})


	r.POST("/process", func(c *gin.Context) {
		var json TaskRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "Processed: " + json.Data,
			"length": len(json.Data),
		})
	})

	r.Run(":8080")
}