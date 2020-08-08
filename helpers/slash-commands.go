package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elvisgastelum/devsearchbot/model"
)

// dataText is used to fill in text from Google API data
var dataText = make(map[string]string)

// HandleMessage is function for handle the incomming messages
func HandleMessage(url, userName, text string) {
	answer := searchAnswer(text)
	response, err := dataBinding(answer).ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	post, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(response))
	if err != nil {
		log.Fatal(err)
	}
	post.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	executePost, err := client.Do(post)
	if err != nil {
		log.Fatal(err)
	}
	defer executePost.Body.Close()
}

func searchAnswer(text string) model.SearchResults {
	text = strings.Replace(text, " ", "+", -1)
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=AIzaSyD8QNzBdjzt3ZNEbGTz4P1rSAnvDPtbrUU&cx=005033773481765961543:gti8czyzyrw&num=3&q=%s", text)
	googleClient := http.Client{
		Timeout: time.Second * 3,
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
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	value := apiMessage(body)
	return value
}

func apiMessage(jsonRaw []byte) model.SearchResults {
	jsonStructure := model.SearchResults{}
	jsonErr := json.Unmarshal(jsonRaw, &jsonStructure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return jsonStructure
}

func dataBinding(data model.SearchResults) *model.SlashCommandResponse {
	if len(data.Items) < 3 {
		return nil
	}

	slashCommandResponse := model.SlashCommandResponse{}
	blocks := make([]map[string]interface{}, 4)

	for i := 0; i < 3; i++ {
		item := data.Items[i]

		buttonValue := fmt.Sprintf("*<%s|%s>*\n>_%s_", item.Link, item.Title, strings.Replace(item.Snippet, "\n", " ", -1))
		dataText[buttonValue] = buttonValue

		blocks[i] = map[string]interface{}{
			"type": "section",
			"accessory": map[string]interface{}{
				"type": "button",
				"text": map[string]interface{}{
					"type":  "plain_text",
					"text":  "Send",
					"emoji": true,
				},
				"value": dataText[buttonValue],
			},
			"text": map[string]interface{}{
				"type": "mrkdwn",
				"text": dataText[buttonValue],
			},
		}

	}

	blocks[3] = map[string]interface{}{
		"type": "actions",
		"elements": []map[string]interface{}{
			{
				"type": "button",
				"text": map[string]interface{}{
					"type":  "plain_text",
					"text":  "Cancel",
					"emoji": true,
				},
				"style": "danger",
				"value": "cancel",
			},
		},
	}

	slashCommandResponse["blocks"] = blocks

	return &slashCommandResponse
}

// ButtonAction determines the response a certain button will give
func ButtonAction(action, responseURL string) {

	var jsonBytes []byte

	if action == "cancel" {
		jsonBytes = []byte(`{"text":null,"response_type":"ephemeral","replace_original":true,"delete_original":true}`)
	} else {
		jsonBytes = []byte(fmt.Sprintf(`{"text":"%s","response_type":"in_channel","replace_original":true,"delete_original":true}`, dataText[action]))
	}

	post, err := http.NewRequest(http.MethodPost, responseURL, bytes.NewBuffer(jsonBytes))
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
