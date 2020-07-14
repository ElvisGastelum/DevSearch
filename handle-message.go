package devsearchbot

import (
	"log"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)


// handleMessage is function for handle the incomming messages
func handleMessage(rtm *slack.RTM, ev *slack.MessageEvent) {
	text := ev.Text
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	log.Printf("User Text: %s\n", text)

	pingPong(rtm, text , ev)

}

// pingPong is a example function for response pong when the user type ping on the chat  
func pingPong (rtm *slack.RTM, text string, ev *slack.MessageEvent){
	info := rtm.GetInfo()

	matched, _ := regexp.MatchString("ping", text)

	if ev.User != info.User.ID && matched {
		rtm.SendMessage(rtm.NewOutgoingMessage("pong", ev.Channel))
	}
}
