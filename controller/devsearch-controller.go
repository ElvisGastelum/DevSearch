package controller

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

type controller struct{}

type DevSearchController interface {
	SlashCommand(response http.ResponseWriter, request *http.Request)
	Actions(response http.ResponseWriter, request *http.Request)
}

func NewDevSearchController() DevSearchController {
	return &controller{}
}

func (*controller) SlashCommand(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	postURL, userName, text := request.PostForm.Get(`response_url`), request.PostForm.Get(`user_name`), request.PostForm.Get(`text`)
	handleMessage(postURL, userName, text)
}

func (*controller) Actions(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	payload := model.Payload{}

	jsonErr := json.Unmarshal([]byte(request.Form.Get("payload")), &payload)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(payload.Actions[0].Value)
}

// handleMessage is function for handle the incomming messages
func handleMessage(URL, userName, text string) {
	getAnswer := searchAnswer(text)
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

func searchAnswer(text string) model.SearchResults {
	text = strings.Replace(text, " ", "+", -1)
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=AIzaSyD8QNzBdjzt3ZNEbGTz4P1rSAnvDPtbrUU&cx=005033773481765961543:gti8czyzyrw&num=3&q=%s", text)
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

func apiMessage(jsonRaw []byte) model.SearchResults {
	jsonStructure := model.SearchResults{}
	jsonErr := json.Unmarshal(jsonRaw, &jsonStructure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return jsonStructure
}

func dataBinding(data model.SearchResults) string {
	slackBlock := `{"blocks":[`
	for i := 0; i < 3; i++ {
		item := data.Items[i]
		slackBlock += fmt.Sprintf(`{"type":"section","text":{"type":"mrkdwn","text":"*<%s|%s>*\n>_%s_"},"accessory":{"type":"button","text":{"type":"plain_text","text":"Send","emoji":true},"value":"button_%d"}},`, item.Link, item.Title, strings.Replace(item.Snippet, "\n", " ", -1), i)
	}
	slackBlock += `{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"Cancel","emoji":true},"style":"danger","value":"click_me_123"}]}`
	slackBlock += `]}`
	return slackBlock
}
