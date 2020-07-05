package slackbot

import (
	"os"

	"github.com/slack-go/slack"
)

var (
	// SlackClient is the client of slack
	SlackClient *slack.Client

	// Rtm is the runtime execution for slack
	Rtm *slack.RTM
)

func init() {
	SlackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	Rtm = SlackClient.NewRTM()
	go Rtm.ManageConnection()

	
}

// HelloName is for return a simple greet
func HelloName(name string) string {
	return "Hello, " + name + "!"
}