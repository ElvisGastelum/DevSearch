package devsearchbot

import (
	"log"
	"net/http"
	"bytes"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
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

// handleMessage is function for handle the incomming messages
func handleMessage(URL, userName string) {
	getAnswer := searchAnswer()
	slackMessage := dataBinding(getAnswer)
	var jsonStr = []byte(slackMessage)
	post, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
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
	jsonStructure := SearchResults{}
	jsonErr := json.Unmarshal(jsonRaw, &jsonStructure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return jsonStructure
}

func dataBinding( data SearchResults) string {
	 slackBlock := `{"blocks":[`
	 for i := 0; i < 3; i++ {
	 	item := data.Items[i]
		slackBlock += fmt.Sprintf(`{"type":"section","text":{"type":"mrkdwn","text":"*<%s|%s>*\n>_%s_"},"accessory":{"type":"button","text":{"type":"plain_text","text":"Send","emoji":true},"value":"click_me_123"}},`, item.Link, item.Title, strings.Replace(item.Snippet, "\n", " ", -1))
	 }
	slackBlock += `{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"Cancel","emoji":true},"style":"danger","value":"click_me_123"}]}`
	slackBlock += `]}` 
	return slackBlock
}