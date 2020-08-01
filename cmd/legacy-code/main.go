package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	//"net/url"
	//"html"

	"github.com/slack-go/slack"
)

var (
	slackClient *slack.Client
)

type SearchResults struct {
	Items []Item `json:items`
}

type Item struct {
	Link    string `json:link`
	Snippet string `json:snippet`
	Title   string `json:title`
}

type TextInfo struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Block struct {
	Type    string   `json:"type"`
	Text    TextInfo `json:"text"`
	BlockID string   `json:"block_id"`
}

type Payload struct {
	Blocks []Block `json:"blocks"`
}

func main() {
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

func handleMessage(URL, userName string) {
	//fmt.Printf("%v\n", ev)
	//structure := searchAnswer()
	//slackMessage := dataBinding(structure)
	//slackClient.PostEphemeral(channelID, userID, slack.MsgOptionBlocks(slackMessage...))
	var jsonStr = []byte(`{"text":"Hello there, ` + fmt.Sprintf("%s", userName) +`!"}`)
	post, err := http.NewRequest("POST",URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	post.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	executePost, er := client.Do(post)
	if er != nil {
		log.Fatal(er)
	}
	defer executePost.Body.Close()
}

func replyToUser(jsonMessage []byte) {
	resp, err := http.Post(os.Getenv("WEB_HOOK"), "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, getErr := ioutil.ReadAll(resp.Body)
	if getErr != nil {
		log.Fatalln(getErr)
	}
	log.Println(body)
}

func searchAnswer() SearchResults {
	url := "https://www.googleapis.com/customsearch/v1?key=AIzaSyD8QNzBdjzt3ZNEbGTz4P1rSAnvDPtbrUU&cx=005033773481765961543:gti8czyzyrw&num=3&q=golang"
	googleClient := http.Client{
		Timeout: time.Second * 3, // Maximum of 3 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Isacc Hernandez")
	res, getErr := googleClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	value := apiMessage(body)
	return value
}

func apiMessage(jsonRaw []byte) SearchResults {
	structure := SearchResults{}
	jsonErr := json.Unmarshal(jsonRaw, &structure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return structure
}

func dataBinding(data SearchResults) []slack.Block {
	var slackBlock []slack.Block
	for i := 0; i < 3; i++ {
		item := data.Items[i]
		btn := slack.NewButtonBlockElement("", "", slack.NewTextBlockObject("plain_text", "Send", false, false))
		textBlock := fmt.Sprintf("*<%s|%s>*\n>_%s_", item.Link, item.Title, strings.Replace(item.Snippet, "\n", " ", -1))
		msgBlock := slack.NewTextBlockObject("mrkdwn", textBlock, false, false)
		slackBlock = append(slackBlock, slack.NewSectionBlock(msgBlock, nil, slack.NewAccessory(btn)))
	}	
	cancelBtn := slack.NewButtonBlockElement("", "", slack.NewTextBlockObject("plain_text", "Cancel", false, false))
	actionBlock := slack.NewActionBlock("", cancelBtn)
	slackBlock = append(slackBlock, actionBlock)
	return slackBlock
}