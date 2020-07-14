package devsearchbot

import (

  "github.com/slack-go/slack"
  "github.com/elvisgastelum/utils"
)

// Bot is the instance of dev search
// Simple example use:
//
// bot := devsearchbot.Bot{}
//
// bot.Start()
type Bot struct{}


// Start the server of dev search
// To run this server, you need set up the enviroment var
// SLACK_ACCESS_TOKEN w/ the token of the slack bot app
func (b *Bot) Start() {
	token := utils.Getenv("SLACK_ACCESS_TOKEN")
  api := slack.New(token)
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  loop(rtm)
}


