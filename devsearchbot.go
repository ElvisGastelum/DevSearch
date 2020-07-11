package devsearchbot

import (
  "log"
  "regexp"
  "strings"

  "github.com/slack-go/slack"
  "github.com/elvisgastelum/devsearchbot/utils"
)

// New create a new instance of dev search bot
func New() {
	token := utils.Getenv("SLACK_ACCESS_TOKEN")
  api := slack.New(token)
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  loop(rtm)
}


// loop is for listen messages events from slack
func loop(rtm *slack.RTM) {
	Loop:
	for {
    select {
    case msg := <-rtm.IncomingEvents:
      log.Printf("Event Received: %v\n", msg.Type)
      switch ev := msg.Data.(type) {

      case *slack.MessageEvent:
        info := rtm.GetInfo()

        text := ev.Text
        text = strings.TrimSpace(text)
        text = strings.ToLower(text)
        log.Printf("User Text: %s\n", text)

        matched, _ := regexp.MatchString("ping", text)

        if ev.User != info.User.ID && matched {
          rtm.SendMessage(rtm.NewOutgoingMessage("pong", ev.Channel))
        }

      case *slack.RTMError:
        log.Printf("Error: %s\n", ev.Error())

      case *slack.InvalidAuthEvent:
        log.Printf("Invalid credentials")
        break Loop

      default:
        // Take no action
      }
    }
  }
}	