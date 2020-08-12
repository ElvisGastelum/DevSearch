package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elvisgastelum/devsearchbot/helpers"
	"github.com/elvisgastelum/devsearchbot/model"
)

var (
	cancelBlock = map[string]interface{}{
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
	acceptBlock = map[string]interface{}{
		"type": "actions",
		"elements": []map[string]interface{}{
			{
				"type": "button",
				"text": map[string]interface{}{
					"type":  "plain_text",
					"text":  "Accept",
					"emoji": true,
				},
				"style": "primary",
				"value": "cancel",
			},
		},
	}
)

// SlashCommands is function for handle the incomming messages
func SlashCommands(url, text string) {
	answer, err := searchAnswer(text)
	if err != nil {
		log.Fatal(err)
	}

	response, err := slashCommandResponse(answer).ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.NewPostRequest(response, url)
	if err != nil {
		log.Fatal(err)
	}
}

func searchAnswer(text string) (model.SearchResults, error) {
	text = strings.Replace(text, " ", "+", -1)
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=AIzaSyD8QNzBdjzt3ZNEbGTz4P1rSAnvDPtbrUU&cx=005033773481765961543:gti8czyzyrw&num=3&q=%s", text)
	var value model.SearchResults

	googleClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return value, err
	}
	req.Header.Set("User-Agent", "Isacc Hernandez")
	res, err := googleClient.Do(req)
	if err != nil {
		return value, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return value, err
	}
	value = apiMessage(body)
	return value, nil

}

func apiMessage(jsonRaw []byte) model.SearchResults {
	jsonStructure := model.SearchResults{}
	jsonErr := json.Unmarshal(jsonRaw, &jsonStructure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return jsonStructure
}

func slashCommandResponse(data model.SearchResults) *model.SlashCommandResponse {
	slashCommandResponse := model.SlashCommandResponse{}
	blocks := make([]map[string]interface{}, 0)
	countResults := len(data.Items)

	if countResults < 1 {
		blocks = append(blocks, createBlock(
			"The search did not produce any results",
		))
		blocks = append(blocks, acceptBlock)
		
		slashCommandResponse["blocks"] = blocks
		return &slashCommandResponse
	}

	Loop:
	for i := 0; i < countResults; i++ {
		item := data.Items[i]

		textValue := fmt.Sprintf("*<%s|%s>*\n>_%s_", item.Link, item.Title, strings.Replace(item.Snippet, "\n", " ", -1))

		blocks = append(blocks, createBlock(textValue))
		
		if i == 3 {
			break Loop
		}
	}
	
	blocks = append(blocks, cancelBlock)
	slashCommandResponse["blocks"] = blocks
	return &slashCommandResponse
}

func createBlock(textValue string) map[string]interface{} {

	result :=  map[string]interface{}{
		"type": "section",
		"text": map[string]interface{}{
			"type": "mrkdwn",
			"text": textValue,
		},
	}

	if textValue != "The search did not produce any results" {
		result["accessory"] = map[string]interface{}{
			"type": "button",
			"text": map[string]interface{}{
				"type":  "plain_text",
				"text":  "Send",
				"emoji": true,
			},
			"value": textValue,
		}
	}

	return result

}


