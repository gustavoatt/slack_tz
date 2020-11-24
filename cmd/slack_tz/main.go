package main

import (
	"net/http"
	"slack_tz/internal/logging"
	"slack_tz/internal/slack"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func handleSlackWebHook(c *gin.Context) {
	var req slack.HookRequest
	if err := c.ShouldBind(&req); err != nil {
		logging.Error(c).
			Err(err).
			Msg("Could not parse slash command webhook")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.String(http.StatusOK, "Hello back!")
}

func main() {
	router := gin.New()
	logging.SetupLogging(router)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/slack_tz_webhook", handleSlackWebHook)

	if err := router.Run(); err != nil {
		log.Printf("Could not start server. Failed with error: %s", err)
	}
}
