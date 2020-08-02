package devsearchbot

import (
	"log"
	"net/http"
//	"github.com/slack-go/slack"
//  "github.com/elvisgastelum/utils"
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
	// token := utils.Getenv("SLACK_ACCESS_TOKEN")
  // api := slack.New(token)
  // rtm := api.NewRTM()
	// go rtm.ManageConnection()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { 
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}
		if request.Method == "POST" {	
				err := request.ParseForm()
				if err != nil {
					log.Fatal(err)
				}
				postURL, userName := request.PostForm.Get(`response_url`), request.PostForm.Get(`user_name`)
				handleMessage(postURL, userName)
		} else {
				http.Error(writer, "Invalid request method.", 405)					
		} 
	})				
	
  log.Println("Listening on :3000...")
  log.Fatal(http.ListenAndServe(":3000", nil))

}