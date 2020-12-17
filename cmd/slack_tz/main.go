package main

import (
	"fmt"
	"net/http"
	"slack_tz/internal/botcommands"
	"slack_tz/internal/logging"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
)

var (
	slackCl = slack.New("xoxb-574908150274-1548687321408-qqNtOn6mB7u53246aEs0QpDB")
)

type hookResponse struct {
	Blocks []slack.Block `json:"blocks"`
}

func buildTimezoneTableString(currentUserID string, users []botcommands.UserInfo, startTime time.Time) (string, error) {
	currentUser, err := slackCl.GetUserInfo(currentUserID)
	if err != nil {
		return "", err
	}

	currentUzerLoc, err := time.LoadLocation(currentUser.TZ)
	if err != nil {
		return "", err
	}

	tableStr := &strings.Builder{}
	tableStr.WriteString("```")

	userTimesTable := tablewriter.NewWriter(tableStr)

	// Set headers
	header := []string{"User"}
	for i := -3; i <= 3; i++ {
		header = append(header, startTime.Add(time.Duration(i)*time.Hour).In(currentUzerLoc).Format(time.Kitchen))
	}

	userTimesTable.SetHeader(header)
	for _, userInfo := range users {
		profile, err := slackCl.GetUserInfo(userInfo.UserID)
		if err != nil {
			return "", fmt.Errorf("could not get user profile: %v", err)
		}

		loc, err := time.LoadLocation(profile.TZ)
		if err != nil {
			return "", fmt.Errorf("could not parse timezone: %v", err)
		}

		row := []string{profile.Profile.DisplayName}
		for i := -3; i <= 3; i++ {
			row = append(row, startTime.Add(time.Duration(i)*time.Hour).In(loc).Format(time.Kitchen))
		}
		userTimesTable.Append(row)
	}

	userTimesTable.Render()
	tableStr.WriteString("```")
	return tableStr.String(), nil
}

func handleSlackWebHook(c *gin.Context) {
	var (
		req slack.SlashCommand
		err error
	)
	if req, err = slack.SlashCommandParse(c.Request); err != nil {
		logging.Error(c).
			Err(err).
			Msg("Could not parse slash command webhook")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// Find all users and their timezones
	tableStr, err := buildTimezoneTableString(
		req.UserID, botcommands.ParseSlackTzCommand(req.Text), time.Now().Truncate(time.Hour))
	if err != nil {
		logging.Error(c).
			Err(err).
			Msg("Could not generate timezone table")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, &hookResponse{
		Blocks: []slack.Block{
			slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, tableStr, false, false), nil, nil),
		},
	})
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
