package main

import (
	"fmt"
	"net/http"
	"slack_tz/internal/slack"

	"github.com/gin-gonic/gin"
)

func handleSlackWebHook(c *gin.Context) {
	var req slack.HookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, req)
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/slack_tz_webhook", handleSlackWebHook)

	if err := router.Run(); err != nil {
		fmt.Printf("Could not start server. Failed with error: %s", err)
	}
}
