package devsearchbot

import (
	"log"
	"net/http"
	"encoding/json"
	
)

// Bot is the instance of dev search
// Simple example use:
//
// bot := devsearchbot.Bot{}
//
// bot.Start()
type Bot struct{}

type ActionBlock struct {
	Action []Actions `json:"actions"`
	URL string `json:"response_url"`
}

type Actions struct {
	Value string `json:"value"`	
}


// Start the server of dev search
// To run this server, you need set up the enviroment var
// SLACK_ACCESS_TOKEN w/ the token of the slack bot app
func (b *Bot) Start() {
	http.HandleFunc("/slack/slash-commands/devz-search", func(writer http.ResponseWriter, request *http.Request) { 
		if request.Method == "POST" {	
				err := request.ParseForm()
				if err != nil {
					log.Fatal(err)
				}
				writer.WriteHeader(http.StatusOK)
				postURL, userName, text := request.PostForm.Get(`response_url`), request.PostForm.Get(`user_name`), request.PostForm.Get(`text`)				
				handleMessage(postURL, userName, text)
		} else {
				http.Error(writer, "Invalid request method.", 405)					
		} 
	})			
	
	http.HandleFunc("/slack/actions/devz-search", func(writer http.ResponseWriter, request *http.Request) { 
		if request.Method == "POST" {	
			err := request.ParseForm()
			if err != nil {
				log.Fatal(err)
			}
			jsonPayload := request.Form.Get("payload")
			jsonStructure := ActionBlock{}
			jsonErr := json.Unmarshal([]byte(jsonPayload), &jsonStructure)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}	
			ButtonAction(jsonStructure.Action[0].Value, jsonStructure.URL)
		} else {
			http.Error(writer, "Invalid request method.", 405)					
		} 
	})
	
  log.Println("Listening on :3000...")
  log.Fatal(http.ListenAndServe(":3000", nil))
}