package main

import (
	"github.com/slack-go/slack"
	"github.com/elvisgastelum/dev-search/slackbot"
)

func main() {
	for msg := range slackbot.Rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(ev.BotID) == 0 {
				go handleMessage(ev)
			}
		}
	}
}

func handleMessage(ev *slack.MessageEvent)  {
	
}
