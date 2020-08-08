package controller

import (
	"log"
	"net/http"

	"github.com/elvisgastelum/devsearchbot/app"
	"github.com/elvisgastelum/devsearchbot/model"
)

type controller struct{}

type DevSearchController interface {
	SlashCommands(response http.ResponseWriter, request *http.Request)
	Actions(response http.ResponseWriter, request *http.Request)
}

func NewDevSearchController() DevSearchController {
	return &controller{}
}

func (*controller) SlashCommands(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New command from: %s", request.Form.Get(`user_name`))

	go app.SlashCommands(
		request.Form.Get(`response_url`),
		request.Form.Get(`text`),
	)
}

func (*controller) Actions(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var payload model.Payload

	err = payload.UnmarshallJSON([]byte(request.Form.Get("payload")))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New action from: %s", payload.User.Name)

	go app.ButtonActions(
		payload.Actions[0].Value,
		payload.ResponseURL,
	)
}
