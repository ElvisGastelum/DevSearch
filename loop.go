package devsearchbot

import (
  "log"
  
	"github.com/slack-go/slack"
)

// loop is for listen messages events from slack
func loop(rtm *slack.RTM) {
	Loop:
	for {
    select {
    case msg := <-rtm.IncomingEvents:
      log.Printf("Event Received: %v\n", msg.Type)

      
      switch ev := msg.Data.(type) {

      case *slack.MessageEvent:
        handleMessage(rtm, ev)

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