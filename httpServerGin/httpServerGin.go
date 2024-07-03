package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getDateTime(c *gin.Context) {
	currentTime := time.Now()
	response := currentTime.Format(time.RFC3339)
	c.JSON(http.StatusOK, gin.H{
		"datetime": response,
	})
}

func main() {
	r := gin.Default()
	r.GET("/datetime", getDateTime)
	r.Run(":8090")
}
