package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elvisgastelum/devsearchbot/model"
	"github.com/elvisgastelum/devsearchbot/helpers"
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
	go helpers.HandleMessage(
		request.Form.Get(`response_url`),
		request.Form.Get(`user_name`),
		request.Form.Get(`text`),
	)
}

func (*controller) Actions(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var payload model.Payload

	jsonErr := json.Unmarshal([]byte(request.Form.Get("payload")), &payload)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	helpers.ButtonAction(payload.Actions[0].Value, payload.ResponseURL)
}
